package tcore

//TRankInfo exported
type TRankInfo struct {
	ID         uint64
	Score      int32
	CustomData interface{}
	Rank       int32

	mIdx int32
}

//TRankSort exported
type TRankSort func(aNewData interface{}, aOldData interface{}) bool

//TRank exported
type TRank struct {
	mRankSize     int32 //排行榜长度
	mCurSize      int32 //当前排行榜长度
	mList         []TRankInfo
	mID2IdxMap    map[uint64]int32
	mSortCallBack TRankSort
}

//Init exported
func (pOwn *TRank) Init(aRankSize int32) bool {
	pOwn.mRankSize = aRankSize
	pOwn.mList = make([]TRankInfo, aRankSize)
	pOwn.mID2IdxMap = make(map[uint64]int32)
	for i := int32(0); i < pOwn.mRankSize; i++ {
		pOwn.mList[i].mIdx = i
		pOwn.mList[i].Rank = i + 1
	}
	return true
}

//SetSortCallBack exported
func (pOwn *TRank) SetSortCallBack(aFunc TRankSort) {
	pOwn.mSortCallBack = aFunc
}

//GetRankByID exported
func (pOwn *TRank) GetRankByID(aID uint64) *TRankInfo {
	nIdx, bOK := pOwn.mID2IdxMap[aID]
	if bOK == false {
		return nil
	}
	return &pOwn.mList[nIdx]
}

//GetAllRank exported
func (pOwn *TRank) GetAllRank() []TRankInfo {
	return pOwn.mList
}

//GetMaxSize exported
func (pOwn *TRank) GetMaxSize() int32 {
	return pOwn.mRankSize
}

//GetCurSize exported
func (pOwn *TRank) GetCurSize() int32 {
	return pOwn.mCurSize
}

//UpdateRank exported
func (pOwn *TRank) UpdateRank(aID uint64, aScore int32, aCustom interface{}) {
	//判断是否在榜上，如果在就直接修改再排序
	nIdx, bOK := pOwn.mID2IdxMap[aID]
	if bOK == true {
		pInfo := &pOwn.mList[nIdx]
		pInfo.Score = aScore
		pInfo.CustomData = aCustom
		pOwn.doRank(nIdx)
		return
	}

	//排行榜没有满,直接放在最后一个
	nIdx = pOwn.mCurSize
	if nIdx < pOwn.mRankSize {
		pOwn.mID2IdxMap[aID] = nIdx
		pInfo := &pOwn.mList[nIdx]
		pInfo.ID = aID
		pInfo.Score = aScore
		pInfo.CustomData = aCustom
		pOwn.mCurSize++
		pOwn.doRank(nIdx)
		return
	}

	//排行榜满了，和最后一个比较是否可以上榜
	nIdx = pOwn.mRankSize - 1
	pLastInfo := &pOwn.mList[nIdx]
	if pOwn.mSortCallBack != nil {
		if pOwn.mSortCallBack(aCustom, pLastInfo.CustomData) == false {
			return
		}
	} else {
		if aScore < pLastInfo.Score {
			return
		}
	}
	delete(pOwn.mID2IdxMap, pLastInfo.ID)
	pLastInfo.ID = aID
	pLastInfo.Score = aScore
	pLastInfo.CustomData = aCustom
	pOwn.mID2IdxMap[aID] = nIdx
	pOwn.doRank(nIdx)
}

func (pOwn *TRank) doRank(aIdx int32) {
	nLastIdx := aIdx - 1
	if nLastIdx < 0 {
		return
	}
	pCurInfo := &pOwn.mList[aIdx]
	pLastInfo := &pOwn.mList[nLastIdx]
	if pOwn.mSortCallBack != nil {
		if pOwn.mSortCallBack(pCurInfo.CustomData, pLastInfo.CustomData) == false {
			return
		}
	} else {
		if pCurInfo.Score <= pLastInfo.Score {
			return
		}
	}
	pOwn.exchange(pCurInfo, pLastInfo)
	pOwn.doRank(nLastIdx)
}

func (pOwn *TRank) exchange(aInfo1 *TRankInfo, aInfo2 *TRankInfo) {
	nIdx1 := aInfo1.mIdx
	nIdx2 := aInfo2.mIdx
	pOwn.mID2IdxMap[aInfo1.ID] = nIdx2
	pOwn.mID2IdxMap[aInfo2.ID] = nIdx1

	aInfo1.ID, aInfo2.ID = aInfo2.ID, aInfo1.ID
	aInfo1.Score, aInfo2.Score = aInfo2.Score, aInfo1.Score
	aInfo1.CustomData, aInfo2.CustomData = aInfo2.CustomData, aInfo1.CustomData
}
