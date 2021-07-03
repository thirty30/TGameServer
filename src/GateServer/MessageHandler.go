//Package gateserver MessageHandler
package gateserver

import (
	tp "TProtocol"
	"encoding/binary"
)

type msgHandler func(aSessionID uint64, aBuffer []byte, aSize uint32)

func (pOwn *GateServer) bindHandler(aMsgID int32, aHandler msgHandler) {
	pOwn.mMsgHandlerMap[aMsgID] = aHandler
}

//注册回调函数
func (pOwn *GateServer) registerHandler() {
	pOwn.bindHandler(tp.LOGIC_2_GATE_TRANSFER_CLIENT, pOwn.handlerSendMsg)
}

func (pOwn *GateServer) handlerSendMsg(aSessionID uint64, aBuffer []byte, aSize uint32) {
	//|--8 sid--|--4 msgid--|--n body--|
	//client sid
	nSessionID := binary.BigEndian.Uint64(aBuffer)
	//msgid
	nMsgID := binary.BigEndian.Uint32(aBuffer[8:])

	var msgHead tp.MessageHead
	msgHead.MsgID = int32(nMsgID)
	msgHead.BodySize = aSize - 12
	msgHead.IsCompressed = 0
	sendBuf := make([]byte, 1024)
	msgHead.Serialize(sendBuf, 9)
	copy(sendBuf[9:], aBuffer[12:])
	pOwn.mNet.Write(nSessionID, sendBuf[:9+msgHead.BodySize])
}
