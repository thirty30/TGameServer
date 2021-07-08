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
	PlayerID    uint64 `bson:"_id"`
	UserID      string
	Renamed     bool
	PlayerName  string
	OfflineTime int64
}
