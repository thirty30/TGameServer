//Package logicserver MessageHandler
package logicserver

import (
	tp "TProtocol"
	"encoding/binary"
)

type msgHandler func(aSessionID uint64, aBuffer []byte, aSize uint32)

func (pOwn *LogicServer) bindHandler(aMsgID int32, aHandler msgHandler) {
	pOwn.mMsgHandlerMap[aMsgID] = aHandler
}

//注册回调函数
func (pOwn *LogicServer) registerHandler() {
	pOwn.bindHandler(tp.GATE_2_LOGIC_REGISTER, pOwn.handlerGateConnected)
	pOwn.bindHandler(tp.GATE_2_LOGIC_TRANSFER_CLIENT, pOwn.handlerRedispathMsg)
	//pOwn.bindHandler(tp.GATE_2_LOGIC_KICK_SESSION, pOwn.handlerKickSession)
	pOwn.bindHandler(tp.C2S_TEST, pOwn.handlerTest)
}

func (pOwn *LogicServer) handlerGateConnected(aSessionID uint64, aBuffer []byte, aSize uint32) {
	pOwn.mGateSessionID = aSessionID
	_LOG(LT_LOG, "GateServer is connected!")
}

func (pOwn *LogicServer) handlerRedispathMsg(aSessionID uint64, aBuffer []byte, aSize uint32) {
	//|--8 sid--|--9 head--|--n body--|
	//client sid
	nClientSessionID := binary.BigEndian.Uint64(aBuffer)

	var msgHead tp.MessageHead
	msgHead.Deserialize(aBuffer[8:], aSize-8)
	nMsgID := msgHead.MsgID

	funcHandler, bExist := pOwn.mMsgHandlerMap[nMsgID]
	if bExist == false || funcHandler == nil {
		return
	}
	funcHandler(nClientSessionID, aBuffer[17:], aSize-17)
}

//测试消息
func (pOwn *LogicServer) handlerTest(aSessionID uint64, aBuffer []byte, aSize uint32) {
	var msgRecv tp.Test
	msgRecv.Deserialize(aBuffer, aSize)

	var msgSend tp.Test
	msgSend.Value0 = false
	msgSend.Value1 = 127
	msgSend.Value2 = 2773
	msgSend.Value3 = 728934792
	msgSend.Value4 = msgRecv.Value4
	msgSend.Value5 = 255
	msgSend.Value6 = 65535
	msgSend.Value7 = 4234444444
	msgSend.Value8 = uint64(msgRecv.Value8)
	msgSend.Value9 = 9.1415
	msgSend.Value10 = 9.14159265358979
	msgSend.Value11 = "i'm genius!!!!"
	msgSend.Value12.Value1 = 78
	msgSend.Value12.Value2 = 192.33
	msgSend.Value12.Value3 = "牛逼下班回家gogogogocccccc!!#4#$%^&*%&%&!!!!!!!!!!!!"

	var t1 tp.T2
	t1.Value1 = 1
	t1.Value2 = 1.1
	t1.Value3 = "abc"
	t1.Value4 = 1
	msgSend.AppendValue13(&t1)
	var t2 tp.T2
	t2.Value1 = 2
	t2.Value2 = 2.2
	t2.Value3 = "def"
	t2.Value4 = 2
	msgSend.AppendValue13(&t2)

	msgSend.AppendValue14(1)
	msgSend.AppendValue14(2)
	msgSend.AppendValue14(3)
	msgSend.AppendValue14(4)

	msgSend.AppendValue15("a")
	msgSend.AppendValue15("回家")
	msgSend.AppendValue15("c")
	msgSend.AppendValue15("逼下班")
}
