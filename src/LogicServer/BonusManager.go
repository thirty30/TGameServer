//Package logicserver BonusManager
package logicserver

import tcore "TCore"

const cMaxProbability uint32 = 1000000

//计算出来的奖励结构
type sObtainItem struct {
	mID  int32
	mNum int32
}

//预加载的奖励数据结构
type sBonusItem struct {
	mItemID      int32
	mMaxNum      int32
	mMinNum      int32
	mProbability uint32
}

type sBonus struct {
	mID   int32
	mType int32
	mList []sBonusItem
}

type sBonusManager struct {
	mBonusMap map[int32]*sBonus
}

func (pOwn *sBonusManager) init() error {
	/*	pOwn.mBonusMap = make(map[int32]*sBonus)
		pRofTable := gServerSingleton.getRofManager().GetBonusTable()
		rowNum := pRofTable.GetRows()
		var i int32
		for i = 0; i < rowNum; i++ {
			pRofItem := pRofTable.GetDataByRow(i)
			pBonus := new(sBonus)
			pBonus.mID = pRofItem.GetID()
			pBonus.mType = pRofItem.GetType()

			str := pRofItem.GetBonus()
			strItems := strings.Split(str, ",")
			pBonus.mList = make([]sBonusItem, len(strItems))
			for j := 0; j < len(strItems); j++ {
				strItem := strings.Split(strItems[j], ":")
				if len(strItem) != 3 {
					return tcore.TNewError("Wrong bonus in row:%d", i)
				}
				var err1, err2, err3 error
				tempID, err1 := strconv.Atoi(strItem[0])
				tempNum := strings.Split(strItem[1], "-")
				tempP, err3 := strconv.Atoi(strItem[2])
				if err1 != nil || err2 != nil || err3 != nil {
					return tcore.TNewError("Wrong bonus item in row:%d", i)
				}

				pBonus.mList[j].mItemID = int32(tempID)
				nMinNum, err4 := strconv.Atoi(tempNum[0])
				nMaxNum, err5 := strconv.Atoi(tempNum[len(tempNum)-1])
				if err4 != nil || err5 != nil {
					return tcore.TNewError("Wrong bonus item in row:%d", i)
				}
				pBonus.mList[j].mMinNum = int32(nMinNum)
				pBonus.mList[j].mMaxNum = int32(nMaxNum)
				pBonus.mList[j].mProbability = uint32(tempP)

				pOwn.mBonusMap[pBonus.mID] = pBonus
			}
		}
	*/
	return nil
}

func (pOwn *sBonusManager) calcBonus(aBounsID int32) []sObtainItem {
	pCurBonus := pOwn.mBonusMap[aBounsID]
	if pCurBonus == nil {
		return nil
	}
	var out []sObtainItem
	if pCurBonus.mType == 1 {
		p := tcore.RandUint32() % cMaxProbability
		var tempP uint32
		for _, v := range pCurBonus.mList {
			tempP += v.mProbability
			if tempP > p {
				nNum := tcore.RandUint32Range(uint32(v.mMinNum), uint32(v.mMaxNum))
				out = append(out, sObtainItem{v.mItemID, int32(nNum)})
				break
			}
			tempP += v.mProbability
		}
	} else if pCurBonus.mType == 2 {
		for _, v := range pCurBonus.mList {
			p := tcore.RandUint32() % cMaxProbability
			if v.mProbability > p {
				nNum := tcore.RandUint32Range(uint32(v.mMinNum), uint32(v.mMaxNum))
				out = append(out, sObtainItem{v.mItemID, int32(nNum)})
			}
		}
	}

	return out
}
