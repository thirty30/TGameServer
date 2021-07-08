package tnet

import "time"

const cAcceptChannelNum = 65536      //新建立链接的channel长度
const cNotifyCloseChannelNum = 65536 //关闭通知的channel长度
const cReceiveMsgChannelNum = 102400 //消息队列容量
const cSessionReadBufLen = 20480     //1024 * 20 //session临时存放读取数据的buffer大小,两次EventDispatch之间存放收到的消息
const cSessionWriteChanLen = 1024    //session临时存放发送数据的channel长度，数组元素是[]byte，两次EventDispatch之间发送消息
const cMaxMessageSize = 2048         //web socket 每个消息最大大小
const cPingPeriod = 50 * time.Second //web socket 服务器发送ping消息时间
const cPongWait = 60 * time.Second   //web socket 服务器等待客户端回复的时间
const cWriteWait = 10 * time.Second  //web socket Time allowed to write a message to the peer.

type onConnect func(aSessionID uint64)                                //连接回调
type onDisconnect func(aSessionID uint64)                             //断开连接回调
type onReceive func(aSessionID uint64, aData []byte, aDataLen uint32) //收到数据回调
type onException func(aSessionID uint64)                              //异常回调
type onGetPacketSize func(aData []byte, aDataLen uint32) uint32       //得到数据包大小

type sMessageObj struct {
	mSessionID uint64
	mData      []byte
	mLen       uint32
}
