package logicserver

//登录加载回调数据
type sLoginInfo struct {
	mSessionID uint64
	mUserID    string
}

//计算出来的奖励结构
type sObtainItem struct {
	mID  int32
	mNum int32
}

const (
	cRedWin          = iota //红队赢
	cBlueWin                //蓝队赢
	cNoDemand               //位阶没有要求
	cContributePoint        //位阶贡献点
	cChallengeNpc           //位阶挑战NPC
	cCompleteTask           //位阶完成任务
	cChallengePlayer        //位阶挑战玩家
)
