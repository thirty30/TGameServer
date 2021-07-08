//Package tcore GUID generator
package tcore

import "math"

//TNumGUIDGenerator 不是并发安全的
type TNumGUIDGenerator struct {
	mInitialNum uint64
	mIDX        uint64
}

func (pOwn *TNumGUIDGenerator) SetInitialNum(aNum uint64) {
	pOwn.mInitialNum = aNum
	pOwn.mIDX = aNum
}

func (pOwn *TNumGUIDGenerator) GetGUID() uint64 {
	pOwn.mIDX++
	if pOwn.mIDX == math.MaxUint64 {
		pOwn.mIDX = pOwn.mInitialNum
	}
	return pOwn.mIDX
}

func (pOwn *TNumGUIDGenerator) Clear() {
	pOwn.mIDX = pOwn.mInitialNum
}

///////////////////////////////////////////////////////////////////////////////////////////

//TTimeGUIDGenerator 不是并发安全的
type TTimeGUIDGenerator struct {
	mIDX        uint32
	mIncSection uint32
	mLastTime   int64
}

func (pOwn *TTimeGUIDGenerator) Init() {
	pOwn.mIncSection = 100000
	pOwn.mLastTime = GetNowTimeStamp()
}

func (pOwn *TTimeGUIDGenerator) SetInitialSection(aDigit int) {
	pOwn.mIncSection = uint32(math.Pow10(int(aDigit)))
}

func (pOwn *TTimeGUIDGenerator) GetGUID() uint64 {
	nowTime := GetNowTimeStamp()
	if nowTime > pOwn.mLastTime {
		pOwn.mLastTime = nowTime
		pOwn.mIDX = 0
	}
	pOwn.mIDX++
	return uint64(nowTime)*uint64(pOwn.mIncSection) + uint64(pOwn.mIDX)
}
