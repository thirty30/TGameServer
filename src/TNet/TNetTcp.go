package tnet

import (
	tp "TProtocol"
	"fmt"
	"net"
	"strconv"
	"time"
)

type sTCPSession struct {
	mSessionID     uint64
	mConn          *net.TCPConn
	mChWrite       chan []byte
	mReadBuffer    []byte
	mReadBufferLen uint32
}

//TCPReactor export
type TCPReactor struct {
	mListenPort         uint16
	mSessionIDGenerater uint64
	mSessionMap         map[uint64]*sTCPSession //连接关系
	mAcceptConnChannel  chan *net.TCPConn       //接收到的连接管道
	mNotifyCloseChannel chan uint64             //关闭连接的通知管道
	mRecvDataChannel    chan *sMessageObj       //分包好的管道

	//callback func
	mCallBackConnect    onConnect
	mCallBackDisconnect onDisconnect
	mCallBackReceive    onReceive
	mCallBackException  onException
}

//Init export
func (pOwn *TCPReactor) Init() {
	pOwn.mListenPort = 0
	pOwn.mSessionIDGenerater = 100
	pOwn.mSessionMap = make(map[uint64]*sTCPSession)
	pOwn.mAcceptConnChannel = make(chan *net.TCPConn, 65536)
	pOwn.mNotifyCloseChannel = make(chan uint64, 65536)
	pOwn.mRecvDataChannel = make(chan *sMessageObj, 102400)
}

//RegisterCallBack export
func (pOwn *TCPReactor) RegisterCallBack(aOnConnected onConnect, aOnDisconnected onDisconnect, aOnReceive onReceive, aOnException onException) {
	pOwn.mCallBackConnect = aOnConnected
	pOwn.mCallBackDisconnect = aOnDisconnected
	pOwn.mCallBackReceive = aOnReceive
	pOwn.mCallBackException = aOnException
}

//Listen export
func (pOwn *TCPReactor) Listen(aPort uint16) error {
	pOwn.mListenPort = aPort
	strAddr := ":" + strconv.Itoa(int(aPort))
	pAddr, err := net.ResolveTCPAddr("tcp", strAddr)
	if err != nil {
		return err
	}
	pListener, err := net.ListenTCP("tcp", pAddr)
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		go pOwn.simulAccpet(pListener)
	}
	return nil
}

//EventDispatch export
func (pOwn *TCPReactor) EventDispatch(aMaxNumPerDeal int, aSecondWait float64) {
	s := time.Now()
	nMsgCount := 0
	select {
	case pMsg := <-pOwn.mRecvDataChannel:
		if pOwn.mCallBackReceive != nil {
			pOwn.mCallBackReceive(pMsg.mSessionID, pMsg.mData, pMsg.mLen)
		}
		nMsgCount++
		if nMsgCount >= aMaxNumPerDeal {
			break
		}
	default:
		if time.Since(s).Seconds() > aSecondWait {
			break
		}
	}

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
		break
	}

}

//Write export
func (pOwn *TCPReactor) Write(aSessionID uint64, aBuffer []byte) {
	pSession, bExist := pOwn.mSessionMap[aSessionID]
	if bExist == false || pSession == nil {
		return
	}
	pSession.mChWrite <- aBuffer
}

//Close export
func (pOwn *TCPReactor) Close(aSessionID uint64, aMilliDelay int) {
	pSession, bExist := pOwn.mSessionMap[aSessionID]
	if bExist == false || pSession == nil {
		return
	}
	go pOwn.simulClose(aSessionID, aMilliDelay)
}

func (pOwn *TCPReactor) newSession(aNewConn *net.TCPConn) *sTCPSession {
	pSession := new(sTCPSession)
	pOwn.mSessionIDGenerater++
	pSession.mSessionID = pOwn.mSessionIDGenerater
	pSession.mConn = aNewConn
	pSession.mChWrite = make(chan []byte, 1024)
	pSession.mReadBuffer = make([]byte, cSessionReadBufLen)
	return pSession
}

func (pOwn *TCPReactor) simulAccpet(aListener *net.TCPListener) {
	for true {
		pConn, err := aListener.AcceptTCP()
		if err != nil {
			continue
		}
		pOwn.mAcceptConnChannel <- pConn
	}
	aListener.Close()
}

func (pOwn *TCPReactor) simulConnectionRead(aSession *sTCPSession) {
	aSession.mConn.SetKeepAlive(true)
	for true {
		nRecvLen, err := aSession.mConn.Read(aSession.mReadBuffer[aSession.mReadBufferLen:])
		if err != nil {
			pOwn.mNotifyCloseChannel <- aSession.mSessionID
			break
		}
		if nRecvLen <= 0 {
			continue
		}
		aSession.mReadBufferLen += uint32(nRecvLen)
		if aSession.mReadBufferLen >= cSessionReadBufLen {
			aSession.mReadBufferLen--
		}

		nLeftLen := aSession.mReadBufferLen
		var nOffset uint32
		for true {
			if nLeftLen <= 0 {
				aSession.mReadBufferLen = 0
				break
			}
			nPackageLen := getPacketSize(aSession.mReadBuffer[nOffset:aSession.mReadBufferLen], nLeftLen)
			if nPackageLen > 0 && nPackageLen <= nLeftLen {
				pMsg := new(sMessageObj)
				pMsg.mSessionID = aSession.mSessionID
				pMsg.mData = make([]byte, nPackageLen)
				copy(pMsg.mData, aSession.mReadBuffer[nOffset:nOffset+nPackageLen])
				pMsg.mLen = nPackageLen
				pOwn.mRecvDataChannel <- pMsg

				nOffset += nPackageLen
				nLeftLen -= nPackageLen

				continue
			}
			if nOffset > 0 {
				tempBuffer := make([]byte, nLeftLen)
				copy(tempBuffer, aSession.mReadBuffer[nOffset:aSession.mReadBufferLen])
				copy(aSession.mReadBuffer, tempBuffer)
				aSession.mReadBufferLen = nLeftLen
				break
			}

		}

	}
}

func (pOwn *TCPReactor) simulConnectionWrite(aSession *sTCPSession) {
	for true {
		pBuffer, chErr := <-aSession.mChWrite
		if chErr == false {
			break
		}
		_, err := aSession.mConn.Write(pBuffer)
		if err != nil {
			break
		}
	}
}

func (pOwn *TCPReactor) simulClose(aSessionID uint64, aMilliDelay int) {
	if aMilliDelay > 0 {
		time.Sleep(time.Millisecond * time.Duration(aMilliDelay))
	}
	pOwn.mNotifyCloseChannel <- aSessionID
}

func getPacketSize(aData []byte, aDataLen uint32) uint32 {
	var msgHead tp.MessageHead
	nHeadLen := msgHead.GetHeadSize()
	if aDataLen < nHeadLen {
		return 0
	}
	msgHead.Deserialize(aData, aDataLen)
	return msgHead.BodySize + nHeadLen
}

//ConnectHost export
func (pOwn *TCPReactor) ConnectHost(aIP string, aPort uint16) error {
	strAddr := fmt.Sprintf("%s:%d", aIP, aPort)
	pAddr, err := net.ResolveTCPAddr("tcp", strAddr)
	if err != nil {
		return err
	}

	pConn, err := net.DialTCP("tcp", nil, pAddr)
	if err != nil {
		return err
	}
	pOwn.mAcceptConnChannel <- pConn
	return nil
}
