package logicserver

import (
	"Rof"
	"TCore"
	tnet "TNet"
	tp "TProtocol"
	"encoding/binary"
	"math/rand"
	"time"
)

const (
	cDBCryptoKey = "thirty1234567890" //数据库加解密密钥
)

var gServerSingleton *LogicServer

//LogicServer export
type LogicServer struct {
	mRun           bool
	mConfig        sConfig
	mNet           tnet.TCPReactor
	mLogManager    sLog
	mRofManager    rof.RofManager
	mBonusManager  sBonusManager
	mDBManager     sDBManager
	mMsgHandlerMap map[int32]msgHandler
	mSessionMgr    sSessionManager

	//Server Session ID
	mGateSessionID uint64

	//debug mode
	mMonitor      *sMonitor
	mDebugGM      *sDebugGM
	mDebugConsole *tcore.DebugConsole
}

//Init export
func (pOwn *LogicServer) Init() bool {
	gServerSingleton = pOwn
	pOwn.mRun = true
	pOwn.mSessionMgr.init()

	var err error
	if pOwn.initConfig() == false {
		return false
	}

	//init log manager
	err = pOwn.mLogManager.init(pOwn.mConfig.LogLevel)
	if err != nil {
		return false
	}

	//init debug mode
	pOwn.initDebugMode()

	//init resource of file
	err = pOwn.mRofManager.Init()
	if err != nil {
		pOwn.mLogManager.error(err.Error())
		return false
	}

	//init db
	pConfig := pOwn.mConfig
	err = pOwn.mDBManager.init(pConfig.DBPath, pConfig.DBPort, pConfig.DBName, pConfig.DBUser, pConfig.DBPwd, cDBCryptoKey)
	if err != nil {
		pOwn.getLogManager().error("Init db manager fail," + err.Error())
		return false
	}

	//init net
	if pOwn.initNet() == false {
		pOwn.getLogManager().error("Init net fail...")
		return false
	}

	//init manager
	if pOwn.initLogicManager() == false {
		return false
	}

	pOwn.getLogManager().log("LogicServer started...")
	return true
}

//Run export
func (pOwn *LogicServer) Run() {
	for pOwn.mRun {
		pOwn.mNet.EventDispatch(100, 0.01)
		pOwn.mDBManager.eventDispatch()
	}
}

//Clear export
func (pOwn *LogicServer) Clear() {
	//pOwn.mDBManager.Clear()
	pOwn.mLogManager.clear()
}

//init http net
func (pOwn *LogicServer) initNet() bool {
	pOwn.mMsgHandlerMap = make(map[int32]msgHandler)
	pOwn.registerHandler()
	pOwn.mNet.Init()
	pOwn.mNet.RegisterCallBack(pOwn.onConnected, pOwn.onDisconnect, pOwn.onReceive, pOwn.onException)
	pOwn.mNet.Listen(uint16(pOwn.mConfig.LogicPort))

	return true
}

func (pOwn *LogicServer) onConnected(aSessionID uint64) {

}

func (pOwn *LogicServer) onDisconnect(aSessionID uint64) {

}

func (pOwn *LogicServer) onReceive(aSessionID uint64, aData []byte, aDataLen uint32) {
	var msgHead tp.MessageHead
	msgHead.Deserialize(aData, aDataLen)
	nMsgID := msgHead.MsgID
	pOwn.mMsgHandlerMap[nMsgID](aSessionID, aData[9:], aDataLen-9)
}

func (pOwn *LogicServer) onException(aSessionID uint64) {

}

func (pOwn *LogicServer) sendMsgToClient(aSessionID uint64, aMsgID int32, aMsg tp.MessageBase) {
	sendBuf := make([]byte, 1024)
	//|--9 head--|--8 sid--|--4 msgid--|--n body--|
	//sid
	binary.BigEndian.PutUint64(sendBuf[9:], aSessionID)
	//msgid
	binary.BigEndian.PutUint32(sendBuf[17:], uint32(aMsgID))
	//body
	nSendBufLen := aMsg.Serialize(sendBuf[21:], 1024)

	var shellHead tp.MessageHead
	shellHead.MsgID = tp.LOGIC_2_GATE_TRANSFER_CLIENT
	shellHead.BodySize = 12 + nSendBufLen
	shellHead.Serialize(sendBuf, 9)

	pOwn.mNet.Write(pOwn.mGateSessionID, sendBuf[:21+nSendBufLen])
}

//init debug mode
func (pOwn *LogicServer) initDebugMode() {
	if pOwn.mConfig.IsDebug == 1 {
		//init debug consol
		pOwn.mDebugConsole = new(tcore.DebugConsole)
		pOwn.mDebugConsole.Init(debugCmdCallBack)
	}
	if pOwn.mConfig.GMEnalbe == 1 {
		//init debug gm
		pOwn.mDebugGM = new(sDebugGM)
	}
	if pOwn.mConfig.MonitorEnable == 1 {
		//init monitor
		pOwn.mMonitor = new(sMonitor)
		pOwn.mMonitor.init()
	}
}

func (pOwn *LogicServer) initLogicManager() bool {
	err := pOwn.getBonusManager().init()
	if err != nil {
		pOwn.getLogManager().error("Init bonus manager fail," + err.Error())
		return false
	}

	return true
}

func (pOwn *LogicServer) getNowTimeStamp() int64 {
	return time.Now().Unix()
}

//随机数
func (pOwn *LogicServer) randUint32() uint32 {
	return rand.Uint32()
}

//随机范围
func (pOwn *LogicServer) randUint32Range(aMin uint32, aMax uint32) uint32 {
	if aMax-aMin == 0 {
		return aMin
	}
	return rand.Uint32()%(aMax-aMin) + aMin
}

func (pOwn *LogicServer) getConfig() *sConfig {
	return &pOwn.mConfig
}

func (pOwn *LogicServer) getRofManager() *rof.RofManager {
	return &pOwn.mRofManager
}

func (pOwn *LogicServer) getDBManager() *sDBManager {
	return &pOwn.mDBManager
}

func (pOwn *LogicServer) getLogManager() *sLog {
	return &pOwn.mLogManager
}

func (pOwn *LogicServer) getMonitor() *sMonitor {
	return pOwn.mMonitor
}

func (pOwn *LogicServer) getDebugGM() *sDebugGM {
	return pOwn.mDebugGM
}

func (pOwn *LogicServer) getDebugConsole() *tcore.DebugConsole {
	return pOwn.mDebugConsole
}

func (pOwn *LogicServer) getBonusManager() *sBonusManager {
	return &pOwn.mBonusManager
}

func (pOwn *LogicServer) getSessionMgr() *sSessionManager {
	return &pOwn.mSessionMgr
}

func (pOwn *LogicServer) getPlayerBySessionID(aSessionID uint64) *sPlayer {
	pSession := pOwn.mSessionMgr.findSessionBySessionID(aSessionID)
	if pSession == nil {
		return nil
	}
	return pSession.mPlayer
}

func (pOwn *LogicServer) getPlayer(aPlayerID uint64) *sPlayer {
	pSession := pOwn.mSessionMgr.findSessionByPlayerID(aPlayerID)
	if pSession == nil {
		return nil
	}
	return pSession.mPlayer
}
