//Package gateserver strcut
package gateserver

import (
	tcore "TCore"
	tnet "TNet"
	tp "TProtocol"
	"encoding/binary"
	"time"
)

type LogType uint8

const (
	LT_DEBUG   LogType = 0
	LT_LOG     LogType = 1
	LT_WARNING LogType = 2
	LT_ERROR   LogType = 3
)

func _LOG(aType LogType, aFormat string, aParms ...interface{}) {
	switch aType {
	case LT_DEBUG:
		gServerSingleton.mLogManager.Debug(aFormat, aParms...)

	case LT_LOG:
		gServerSingleton.mLogManager.Log(aFormat, aParms...)

	case LT_WARNING:
		gServerSingleton.mLogManager.Warning(aFormat, aParms...)

	case LT_ERROR:
		gServerSingleton.mLogManager.Error(aFormat, aParms...)
	}
}

func getMessagePacketSize(aData []byte, aDataLen uint32) uint32 {
	var msgHead tp.MessageHead
	nHeadLen := msgHead.GetHeadSize()
	if aDataLen < nHeadLen {
		return 0
	}
	msgHead.Deserialize(aData, aDataLen)
	return msgHead.BodySize + nHeadLen
}

var gServerSingleton *GateServer

//GateServer exported
type GateServer struct {
	mRun           bool
	mConfig        sConfig
	mLogManager    tcore.TLog
	mNet           tnet.TCPReactor
	mMsgHandlerMap map[int32]msgHandler

	mLogicSessionID uint64
}

//Init exported
func (pOwn *GateServer) Init() bool {
	gServerSingleton = pOwn
	pOwn.mRun = true
	if pOwn.initConfig() == false {
		return false
	}

	//init log manager
	err := pOwn.mLogManager.Init("./log/GateServer", pOwn.mConfig.LogLevel, (pOwn.mConfig.IsDebug == 1))
	if err != nil {
		return false
	}

	//init net
	pOwn.mMsgHandlerMap = make(map[int32]msgHandler)
	pOwn.registerHandler()
	pOwn.mNet.Init()
	pOwn.mNet.RegisterCallBack(pOwn.onConnected, pOwn.onDisconnect, pOwn.onReceive, pOwn.onException, getMessagePacketSize)
	err = pOwn.mNet.ConnectHost("127.0.0.1", uint16(pOwn.mConfig.LogicPort))
	if err != nil {
		_LOG(LT_ERROR, err.Error())
		return false
	}

	_LOG(LT_LOG, "GateServer started...")
	return true
}

//Run exported
func (pOwn *GateServer) Run() {
	for pOwn.mRun {
		st := time.Now()

		pOwn.mNet.EventDispatch(1000)
		time.Sleep(time.Millisecond * 10)

		lt := 20 - time.Since(st).Milliseconds() //服务器每秒50帧
		if lt > 0 {
			time.Sleep(time.Duration(lt) * time.Millisecond)
		}
	}
}

//Clear exported
func (pOwn *GateServer) Clear() {
	pOwn.mLogManager.Clear()
}

func (pOwn *GateServer) onConnected(aSessionID uint64) {
	if pOwn.mLogicSessionID == 0 {
		pOwn.mLogicSessionID = aSessionID

		//告诉逻辑服务器连上了
		var msgSend tp.CommonN8
		pOwn.sendMsgToLogic(aSessionID, tp.GATE_2_LOGIC_REGISTER, &msgSend)

		//监听外部端口
		err := pOwn.mNet.Listen(uint16(pOwn.mConfig.ExternalPort))
		if err != nil {
			_LOG(LT_ERROR, err.Error())
		}
	}
}

func (pOwn *GateServer) onDisconnect(aSessionID uint64) {
	if aSessionID != pOwn.mLogicSessionID {
		var msgSend tp.CommonU64
		msgSend.Value = aSessionID
		pOwn.sendMsgToLogic(aSessionID, tp.GATE_2_LOGIC_KICK_SESSION, &msgSend)
	}
}

func (pOwn *GateServer) onReceive(aSessionID uint64, aData []byte, aDataLen uint32) {
	var msgHead tp.MessageHead
	msgHead.Deserialize(aData, aDataLen)
	nMsgID := msgHead.MsgID
	if aSessionID == pOwn.mLogicSessionID {
		pOwn.mMsgHandlerMap[nMsgID](aSessionID, aData[9:aDataLen], aDataLen-9)
	} else {
		pOwn.transformToLogic(aSessionID, aData, aDataLen)
	}
}

func (pOwn *GateServer) onException(aSessionID uint64) {
}

func (pOwn *GateServer) transformToLogic(aSessionID uint64, aBuffer []byte, aSize uint32) {
	sendBuf := make([]byte, 1024)
	//|--9 head--|--8 sid--|--9 head--|--n body--|
	//sid
	binary.BigEndian.PutUint64(sendBuf[9:], aSessionID)
	//real msg
	copy(sendBuf[17:], aBuffer)

	var shellHead tp.MessageHead
	shellHead.MsgID = tp.GATE_2_LOGIC_TRANSFER_CLIENT
	shellHead.BodySize = 8 + aSize
	shellHead.Serialize(sendBuf, 9)

	pOwn.mNet.Write(pOwn.mLogicSessionID, sendBuf[:17+aSize])
}

func (pOwn *GateServer) sendMsgToLogic(aSessionID uint64, aMsgID int32, aMsg tp.MessageBase) {
	sendBuf := make([]byte, 1024)
	var msgHead tp.MessageHead
	nSendBufLen := aMsg.Serialize(sendBuf[9:], 1024)

	msgHead.MsgID = aMsgID
	msgHead.BodySize = nSendBufLen
	msgHead.Serialize(sendBuf[:9], 9)

	pOwn.mNet.Write(pOwn.mLogicSessionID, sendBuf[:nSendBufLen+9])
}

func (pOwn *GateServer) getConfig() *sConfig {
	return &pOwn.mConfig
}
