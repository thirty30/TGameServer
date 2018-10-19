package logicserver

//数据库表名
const cDBTablePlayer = "Player" //玩家表
const cDBTableUser = "User"     //用户表

//用户信息
type sDBUser struct {
	UserID        string `bson:"_id"`
	PlayerID      uint64
	RegisterTime  int64
	LastLoginTime int64
}

//玩家数据
type sDBPlayer struct {
	PlayerID       uint64 `bson:"_id"`
	UserID         string
	Renamed        bool
	PlayerName     string
	Level          int32
	Exp            uint64
	QiYun          int32
	Age            uint64 //年龄 单位时辰
	OfflineTime    int64
	XiuLianEndTime int64 //修炼结束时间
	XiuLianGetTime int64 //灵气领取时间
	XiuLianItemID  int32 //修炼物品
	XiuLianLingQi  int32 //修炼获得的总灵气
	SelectFaction  int32 //门派ID
	HierarchyLv    int32 //位阶等级
}
