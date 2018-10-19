package tnet

import "time"

const cMaxConnectionNum = 20480 //最大连接数
const cMsgQueueNum = 10240      //消息队列容量
const cConnReadBufLen = 1024 * 10
const cSessionReadBufLen = cConnReadBufLen * 2 //session临时存放数据的buffer大小
const cMaxMessageSize = 2048                   //每个消息最大大小
const cPingPeriod = 50 * time.Second           //服务器发送ping消息时间
const cPongWait = 60 * time.Second             //服务器等待客户端回复的时间
const cWriteWait = 10 * time.Second            // Time allowed to write a message to the peer.

type onConnect func(aSessionID uint64)                                //连接回调
type onDisconnect func(aSessionID uint64)                             //断开连接回调
type onReceive func(aSessionID uint64, aData []byte, aDataLen uint32) //收到数据回调
type onException func(aSessionID uint64)                              //异常回调

type sMessageObj struct {
	mSessionID uint64
	mData      []byte
	mLen       uint32
}
