package tnet

import (
	"net/http"
	"strconv"
	"time"

	ws "github.com/gorilla/websocket"
)

type sWSSession struct {
	mSessionID  uint64
	mConn       *ws.Conn
	mChWrite    chan []byte
	mReadBuffer []byte
}

//WSReactor export
type WSReactor struct {
	mListenPort         uint16
	mSessionIDGenerater uint64
	mSessionMap         map[uint64]*sWSSession //连接关系
	mAcceptConnChannel  chan *ws.Conn          //接收到的连接管道
	mNotifyCloseChannel chan uint64            //关闭连接的通知管道
	mRecvDataChannel    chan *sMessageObj      //分包好的管道

	//callback func
	mCallBackConnect    onConnect
	mCallBackDisconnect onDisconnect
	mCallBackReceive    onReceive
	mCallBackException  onException
}

//Init export
func (pOwn *WSReactor) Init() {
	pOwn.mListenPort = 0
	pOwn.mSessionIDGenerater = 1
	pOwn.mSessionMap = make(map[uint64]*sWSSession)
	pOwn.mAcceptConnChannel = make(chan *ws.Conn, cAcceptChannelNum)
	pOwn.mNotifyCloseChannel = make(chan uint64, cNotifyCloseChannelNum)
	pOwn.mRecvDataChannel = make(chan *sMessageObj, cReceiveMsgChannelNum)
}

//RegisterCallBack export
func (pOwn *WSReactor) RegisterCallBack(aOnConnected onConnect, aOnDisconnected onDisconnect, aOnReceive onReceive, aOnException onException) {
	pOwn.mCallBackConnect = aOnConnected
	pOwn.mCallBackDisconnect = aOnDisconnected
	pOwn.mCallBackReceive = aOnReceive
	pOwn.mCallBackException = aOnException
}

//Listen export
func (pOwn *WSReactor) Listen(aPort uint16) {
	pOwn.mListenPort = aPort
	go pOwn.simulAccpet()
}

func (pOwn *WSReactor) simulAccpet() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pOwn.registerSession)
	strAddr := ":" + strconv.Itoa(int(pOwn.mListenPort))
	err := http.ListenAndServe(strAddr, mux)
	if err != nil && pOwn.mCallBackException != nil {
		pOwn.mCallBackException(0)
	}
}

func (pOwn *WSReactor) registerSession(aWriter http.ResponseWriter, aReq *http.Request) {
	pUpgrader := new(ws.Upgrader)
	pUpgrader.ReadBufferSize = cMaxMessageSize
	pUpgrader.WriteBufferSize = cMaxMessageSize
	pUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	pConn, err := pUpgrader.Upgrade(aWriter, aReq, nil)
	if err != nil {
		return
	}
	pOwn.mAcceptConnChannel <- pConn
}

//EventDispatch export
//aMaxNumPerDeal 每次处理消息的数量 负数代表处理完所有消息
func (pOwn *WSReactor) EventDispatch(aMaxNumPerDeal int32) {
	bDeal := true
	for bDeal == true {
		select {
		case nCloseSessionID := <-pOwn.mNotifyCloseChannel:
			pSession, bExist := pOwn.mSessionMap[nCloseSessionID]
			if bExist == false || pSession == nil {
				break
			}
			pSession.mConn.Close()
			close(pSession.mChWrite)
			nSID := pSession.mSessionID
			delete(pOwn.mSessionMap, nSID)
			if pOwn.mCallBackDisconnect != nil {
				pOwn.mCallBackDisconnect(nSID)
			}

		case pNewConn := <-pOwn.mAcceptConnChannel:
			pSession := pOwn.newSession(pNewConn)
			pOwn.mSessionMap[pSession.mSessionID] = pSession
			go pOwn.simulConnectionRead(pSession)
			go pOwn.simulConnectionWrite(pSession)
			if pOwn.mCallBackConnect != nil {
				pOwn.mCallBackConnect(pSession.mSessionID)
			}

		default:
			bDeal = false
		}
	}

	//读数据
	nMsgCount := int32(0)
	bDeal = true
	for bDeal == true {
		select {
		case pMsg := <-pOwn.mRecvDataChannel:
			if pOwn.mCallBackReceive != nil {
				pOwn.mCallBackReceive(pMsg.mSessionID, pMsg.mData, pMsg.mLen)
			}
			nMsgCount++
			if aMaxNumPerDeal > 0 && nMsgCount >= aMaxNumPerDeal {
				bDeal = false
				break
			}
			continue

		default:
			bDeal = false

		}
	}
}

//Write export
func (pOwn *WSReactor) Write(aSessionID uint64, aBuffer []byte) {
	pSession, bExist := pOwn.mSessionMap[aSessionID]
	if bExist == false || pSession == nil {
		return
	}
	pSession.mChWrite <- aBuffer
}

//Close export
func (pOwn *WSReactor) Close(aSessionID uint64, aMilliDelay int) {
	pSession, bExist := pOwn.mSessionMap[aSessionID]
	if bExist == false || pSession == nil {
		return
	}
	go pOwn.simulClose(aSessionID, aMilliDelay)
}

func (pOwn *WSReactor) newSession(aNewConn *ws.Conn) *sWSSession {
	pSession := new(sWSSession)
	pOwn.mSessionIDGenerater++
	pSession.mSessionID = pOwn.mSessionIDGenerater
	pSession.mConn = aNewConn
	pSession.mChWrite = make(chan []byte, cSessionWriteChanLen)
	pSession.mReadBuffer = make([]byte, cSessionReadBufLen)
	return pSession
}

func (pOwn *WSReactor) simulConnectionRead(aSession *sWSSession) {
	aSession.mConn.SetReadLimit(cMaxMessageSize)
	aSession.mConn.SetReadDeadline(time.Now().Add(cPongWait))
	aSession.mConn.SetPongHandler(
		func(aAppData string) error {
			aSession.mConn.SetReadDeadline(time.Now().Add(cPongWait))
			return nil
		})
	for {
		_, pBuffer, err := aSession.mConn.ReadMessage()
		if err != nil {
			pOwn.mNotifyCloseChannel <- aSession.mSessionID
			break
		}

		pMsg := new(sMessageObj)
		pMsg.mSessionID = aSession.mSessionID
		pMsg.mData = pBuffer
		pMsg.mLen = uint32(len(pBuffer))
		pOwn.mRecvDataChannel <- pMsg
	}
}

func (pOwn *WSReactor) simulConnectionWrite(aSession *sWSSession) {
	ticker := time.NewTicker(cPingPeriod)
	defer ticker.Stop()
	for {
		select {
		case pBuffer, chErr := <-aSession.mChWrite:
			if chErr == false {
				return
			}
			aSession.mConn.SetWriteDeadline(time.Now().Add(cWriteWait))
			writer, err := aSession.mConn.NextWriter(ws.BinaryMessage)
			if err != nil {
				return
			}
			_, err = writer.Write(pBuffer)
			if err != nil {
				return
			}
			err = writer.Close()
			if err != nil {
				return
			}
		case <-ticker.C:
			aSession.mConn.SetWriteDeadline(time.Now().Add(cWriteWait))
			err := aSession.mConn.WriteMessage(ws.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}

func (pOwn *WSReactor) simulClose(aSessionID uint64, aMilliDelay int) {
	if aMilliDelay > 0 {
		time.Sleep(time.Millisecond * time.Duration(aMilliDelay))
	}
	pOwn.mNotifyCloseChannel <- aSessionID
}
