package logicserver

type sSession struct {
	mSessionID uint64
	mUserID    string
	mPlayer    *sPlayer
}

type sSessionManager struct {
	mSessionIDMap map[uint64]*sSession
	mUserIDMap    map[string]*sSession
	mPlayerIDMap  map[uint64]*sSession
}

func (pOwn *sSessionManager) init() {
	pOwn.mSessionIDMap = make(map[uint64]*sSession)
	pOwn.mUserIDMap = make(map[string]*sSession)
	pOwn.mPlayerIDMap = make(map[uint64]*sSession)
}

func (pOwn *sSessionManager) addNewSession(aSessionID uint64) bool {
	_, bExist := pOwn.mSessionIDMap[aSessionID]
	if bExist == true {
		return false
	}
	pSession := new(sSession)
	pSession.mSessionID = aSessionID
	pOwn.mSessionIDMap[aSessionID] = pSession
	return true
}

func (pOwn *sSessionManager) relateUser(aUserID string, aSession *sSession) bool {
	pOwn.mUserIDMap[aUserID] = aSession
	return true
}

func (pOwn *sSessionManager) relatePlayer(aPlayerID uint64, aSession *sSession) bool {
	_, bExist := pOwn.mPlayerIDMap[aPlayerID]
	if bExist == true {
		return false
	}
	pOwn.mPlayerIDMap[aPlayerID] = aSession
	return true
}

func (pOwn *sSessionManager) findSessionBySessionID(aSessionID uint64) *sSession {
	return pOwn.mSessionIDMap[aSessionID]
}

func (pOwn *sSessionManager) findSessionByUserID(aUserID string) *sSession {
	return pOwn.mUserIDMap[aUserID]
}

func (pOwn *sSessionManager) findSessionByPlayerID(aPlayerID uint64) *sSession {
	return pOwn.mPlayerIDMap[aPlayerID]
}

func (pOwn *sSessionManager) deleteSession(aSessionID uint64) bool {
	pSession, bExist := pOwn.mSessionIDMap[aSessionID]
	if bExist == false {
		return false
	}
	delete(pOwn.mSessionIDMap, aSessionID)

	strUID := pSession.mUserID
	delete(pOwn.mUserIDMap, strUID)

	if pSession.mPlayer == nil {
		return false
	}
	//pSession.mPlayer.offline()
	//delete(pOwn.mPlayerIDMap, pSession.mPlayer.getPlayerID())
	return true
}

func (pOwn *sSessionManager) GetAllOnlinePlayer() []sPlayer {
	var res []sPlayer
	if pOwn.mPlayerIDMap == nil {
		return nil
	}
	for _, value := range pOwn.mPlayerIDMap {
		player := gServerSingleton.getPlayerBySessionID(value.mSessionID)
		if player != nil {
			res = append(res, *player)
		}
	}
	return res
}
