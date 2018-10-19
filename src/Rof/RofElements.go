package rof
import (
"encoding/binary"
"math"
)
type sRofAvatarRow struct {
mID int32
mAvatarName string
mABPath string
mAssetName string
}
func (pOwn *sRofAvatarRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nAvatarNameLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mAvatarName = string(aBuffer[nOffset:nOffset+nAvatarNameLen])
nOffset+=nAvatarNameLen
nABPathLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mABPath = string(aBuffer[nOffset:nOffset+nABPathLen])
nOffset+=nABPathLen
nAssetNameLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mAssetName = string(aBuffer[nOffset:nOffset+nAssetNameLen])
nOffset+=nAssetNameLen
return nOffset
}
func (pOwn *sRofAvatarRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofAvatarRow) GetAvatarName() string { return pOwn.mAvatarName } 
func (pOwn *sRofAvatarRow) GetABPath() string { return pOwn.mABPath } 
func (pOwn *sRofAvatarRow) GetAssetName() string { return pOwn.mAssetName } 
type sRofAvatarTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofAvatarRow
mRowMap map[int32]int32
}
func (pOwn *sRofAvatarTable) newTypeObj() iRofRow {return new(sRofAvatarRow)}
func (pOwn *sRofAvatarTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofAvatarTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofAvatarTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofAvatarRow)}
func (pOwn *sRofAvatarTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofAvatarTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofAvatarRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofAvatarTable) GetDataByID(aID int32) *sRofAvatarRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofAvatarTable) GetDataByRow(aIndex int32) *sRofAvatarRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofAvatarTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofAvatarTable) GetCols() int32 {return pOwn.mColNum}
type sRofBonusRow struct {
mID int32
mType int32
mBonus string
}
func (pOwn *sRofBonusRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mType = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nBonusLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mBonus = string(aBuffer[nOffset:nOffset+nBonusLen])
nOffset+=nBonusLen
return nOffset
}
func (pOwn *sRofBonusRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofBonusRow) GetType() int32 { return pOwn.mType } 
func (pOwn *sRofBonusRow) GetBonus() string { return pOwn.mBonus } 
type sRofBonusTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofBonusRow
mRowMap map[int32]int32
}
func (pOwn *sRofBonusTable) newTypeObj() iRofRow {return new(sRofBonusRow)}
func (pOwn *sRofBonusTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofBonusTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofBonusTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofBonusRow)}
func (pOwn *sRofBonusTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofBonusTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofBonusRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofBonusTable) GetDataByID(aID int32) *sRofBonusRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofBonusTable) GetDataByRow(aIndex int32) *sRofBonusRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofBonusTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofBonusTable) GetCols() int32 {return pOwn.mColNum}
type sRofExampleRow struct {
mID int32
mValue1 int32
mValue2 int64
mValue3 float32
mValue4 float64
mValue5 string
mValue6 int32
}
func (pOwn *sRofExampleRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mValue1 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mValue2 = int64(binary.BigEndian.Uint64(aBuffer[nOffset:]))
nOffset+=8
pOwn.mValue3 = math.Float32frombits(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mValue4 = math.Float64frombits(binary.BigEndian.Uint64(aBuffer[nOffset:]))
nOffset+=8
nValue5Len := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mValue5 = string(aBuffer[nOffset:nOffset+nValue5Len])
nOffset+=nValue5Len
pOwn.mValue6 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
return nOffset
}
func (pOwn *sRofExampleRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofExampleRow) GetValue1() int32 { return pOwn.mValue1 } 
func (pOwn *sRofExampleRow) GetValue2() int64 { return pOwn.mValue2 } 
func (pOwn *sRofExampleRow) GetValue3() float32 { return pOwn.mValue3 } 
func (pOwn *sRofExampleRow) GetValue4() float64 { return pOwn.mValue4 } 
func (pOwn *sRofExampleRow) GetValue5() string { return pOwn.mValue5 } 
func (pOwn *sRofExampleRow) GetValue6() int32 { return pOwn.mValue6 } 
type sRofExampleTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofExampleRow
mRowMap map[int32]int32
}
func (pOwn *sRofExampleTable) newTypeObj() iRofRow {return new(sRofExampleRow)}
func (pOwn *sRofExampleTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofExampleTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofExampleTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofExampleRow)}
func (pOwn *sRofExampleTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofExampleTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofExampleRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofExampleTable) GetDataByID(aID int32) *sRofExampleRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofExampleTable) GetDataByRow(aIndex int32) *sRofExampleRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofExampleTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofExampleTable) GetCols() int32 {return pOwn.mColNum}
type sRofFirstNameRow struct {
mID int32
mFirstName string
}
func (pOwn *sRofFirstNameRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nFirstNameLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mFirstName = string(aBuffer[nOffset:nOffset+nFirstNameLen])
nOffset+=nFirstNameLen
return nOffset
}
func (pOwn *sRofFirstNameRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofFirstNameRow) GetFirstName() string { return pOwn.mFirstName } 
type sRofFirstNameTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofFirstNameRow
mRowMap map[int32]int32
}
func (pOwn *sRofFirstNameTable) newTypeObj() iRofRow {return new(sRofFirstNameRow)}
func (pOwn *sRofFirstNameTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofFirstNameTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofFirstNameTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofFirstNameRow)}
func (pOwn *sRofFirstNameTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofFirstNameTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofFirstNameRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofFirstNameTable) GetDataByID(aID int32) *sRofFirstNameRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofFirstNameTable) GetDataByRow(aIndex int32) *sRofFirstNameRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofFirstNameTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofFirstNameTable) GetCols() int32 {return pOwn.mColNum}
type sRofHeroRow struct {
mID int32
mName string
mAvatarID int32
mSkillBehaviour int32
mSkill string
mScale float32
}
func (pOwn *sRofHeroRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nNameLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mName = string(aBuffer[nOffset:nOffset+nNameLen])
nOffset+=nNameLen
pOwn.mAvatarID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mSkillBehaviour = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nSkillLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mSkill = string(aBuffer[nOffset:nOffset+nSkillLen])
nOffset+=nSkillLen
pOwn.mScale = math.Float32frombits(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
return nOffset
}
func (pOwn *sRofHeroRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofHeroRow) GetName() string { return pOwn.mName } 
func (pOwn *sRofHeroRow) GetAvatarID() int32 { return pOwn.mAvatarID } 
func (pOwn *sRofHeroRow) GetSkillBehaviour() int32 { return pOwn.mSkillBehaviour } 
func (pOwn *sRofHeroRow) GetSkill() string { return pOwn.mSkill } 
func (pOwn *sRofHeroRow) GetScale() float32 { return pOwn.mScale } 
type sRofHeroTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofHeroRow
mRowMap map[int32]int32
}
func (pOwn *sRofHeroTable) newTypeObj() iRofRow {return new(sRofHeroRow)}
func (pOwn *sRofHeroTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofHeroTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofHeroTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofHeroRow)}
func (pOwn *sRofHeroTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofHeroTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofHeroRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofHeroTable) GetDataByID(aID int32) *sRofHeroRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofHeroTable) GetDataByRow(aIndex int32) *sRofHeroRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofHeroTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofHeroTable) GetCols() int32 {return pOwn.mColNum}
type sRofHierarchyRow struct {
mID int32
mHierarchyName string
mHierarchyNum int32
mLimit string
}
func (pOwn *sRofHierarchyRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nHierarchyNameLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mHierarchyName = string(aBuffer[nOffset:nOffset+nHierarchyNameLen])
nOffset+=nHierarchyNameLen
pOwn.mHierarchyNum = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nLimitLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mLimit = string(aBuffer[nOffset:nOffset+nLimitLen])
nOffset+=nLimitLen
return nOffset
}
func (pOwn *sRofHierarchyRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofHierarchyRow) GetHierarchyName() string { return pOwn.mHierarchyName } 
func (pOwn *sRofHierarchyRow) GetHierarchyNum() int32 { return pOwn.mHierarchyNum } 
func (pOwn *sRofHierarchyRow) GetLimit() string { return pOwn.mLimit } 
type sRofHierarchyTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofHierarchyRow
mRowMap map[int32]int32
}
func (pOwn *sRofHierarchyTable) newTypeObj() iRofRow {return new(sRofHierarchyRow)}
func (pOwn *sRofHierarchyTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofHierarchyTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofHierarchyTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofHierarchyRow)}
func (pOwn *sRofHierarchyTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofHierarchyTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofHierarchyRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofHierarchyTable) GetDataByID(aID int32) *sRofHierarchyRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofHierarchyTable) GetDataByRow(aIndex int32) *sRofHierarchyRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofHierarchyTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofHierarchyTable) GetCols() int32 {return pOwn.mColNum}
type sRofItemRow struct {
mID int32
mName int32
mDesc int32
mType int32
mIcon string
mStack int32
mCanUse int32
mAction int32
mParam1 int32
mParam2 int32
mParam3 int32
}
func (pOwn *sRofItemRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mName = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mDesc = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mType = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nIconLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mIcon = string(aBuffer[nOffset:nOffset+nIconLen])
nOffset+=nIconLen
pOwn.mStack = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mCanUse = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mAction = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mParam1 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mParam2 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mParam3 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
return nOffset
}
func (pOwn *sRofItemRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofItemRow) GetName() int32 { return pOwn.mName } 
func (pOwn *sRofItemRow) GetDesc() int32 { return pOwn.mDesc } 
func (pOwn *sRofItemRow) GetType() int32 { return pOwn.mType } 
func (pOwn *sRofItemRow) GetIcon() string { return pOwn.mIcon } 
func (pOwn *sRofItemRow) GetStack() int32 { return pOwn.mStack } 
func (pOwn *sRofItemRow) GetCanUse() int32 { return pOwn.mCanUse } 
func (pOwn *sRofItemRow) GetAction() int32 { return pOwn.mAction } 
func (pOwn *sRofItemRow) GetParam1() int32 { return pOwn.mParam1 } 
func (pOwn *sRofItemRow) GetParam2() int32 { return pOwn.mParam2 } 
func (pOwn *sRofItemRow) GetParam3() int32 { return pOwn.mParam3 } 
type sRofItemTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofItemRow
mRowMap map[int32]int32
}
func (pOwn *sRofItemTable) newTypeObj() iRofRow {return new(sRofItemRow)}
func (pOwn *sRofItemTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofItemTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofItemTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofItemRow)}
func (pOwn *sRofItemTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofItemTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofItemRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofItemTable) GetDataByID(aID int32) *sRofItemRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofItemTable) GetDataByRow(aIndex int32) *sRofItemRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofItemTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofItemTable) GetCols() int32 {return pOwn.mColNum}
type sRofLanguageRow struct {
mID int32
mChinese string
mEnglish string
}
func (pOwn *sRofLanguageRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nChineseLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mChinese = string(aBuffer[nOffset:nOffset+nChineseLen])
nOffset+=nChineseLen
nEnglishLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mEnglish = string(aBuffer[nOffset:nOffset+nEnglishLen])
nOffset+=nEnglishLen
return nOffset
}
func (pOwn *sRofLanguageRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofLanguageRow) GetChinese() string { return pOwn.mChinese } 
func (pOwn *sRofLanguageRow) GetEnglish() string { return pOwn.mEnglish } 
type sRofLanguageTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofLanguageRow
mRowMap map[int32]int32
}
func (pOwn *sRofLanguageTable) newTypeObj() iRofRow {return new(sRofLanguageRow)}
func (pOwn *sRofLanguageTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofLanguageTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofLanguageTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofLanguageRow)}
func (pOwn *sRofLanguageTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofLanguageTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofLanguageRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofLanguageTable) GetDataByID(aID int32) *sRofLanguageRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofLanguageTable) GetDataByRow(aIndex int32) *sRofLanguageRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofLanguageTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofLanguageTable) GetCols() int32 {return pOwn.mColNum}
type sRofLevelRow struct {
mID int32
mName int32
mExp int64
mQiYun int32
mThunder int32
mHeartDevil int32
mLiLiang int32
mGenGu int32
mNaiLi int32
mQiHai int32
mShenFa int32
mSuDu int32
mShangHaiJiaShen float32
mShangHaiJianMian float32
}
func (pOwn *sRofLevelRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mName = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mExp = int64(binary.BigEndian.Uint64(aBuffer[nOffset:]))
nOffset+=8
pOwn.mQiYun = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mThunder = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mHeartDevil = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mLiLiang = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mGenGu = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mNaiLi = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mQiHai = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mShenFa = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mSuDu = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mShangHaiJiaShen = math.Float32frombits(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mShangHaiJianMian = math.Float32frombits(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
return nOffset
}
func (pOwn *sRofLevelRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofLevelRow) GetName() int32 { return pOwn.mName } 
func (pOwn *sRofLevelRow) GetExp() int64 { return pOwn.mExp } 
func (pOwn *sRofLevelRow) GetQiYun() int32 { return pOwn.mQiYun } 
func (pOwn *sRofLevelRow) GetThunder() int32 { return pOwn.mThunder } 
func (pOwn *sRofLevelRow) GetHeartDevil() int32 { return pOwn.mHeartDevil } 
func (pOwn *sRofLevelRow) GetLiLiang() int32 { return pOwn.mLiLiang } 
func (pOwn *sRofLevelRow) GetGenGu() int32 { return pOwn.mGenGu } 
func (pOwn *sRofLevelRow) GetNaiLi() int32 { return pOwn.mNaiLi } 
func (pOwn *sRofLevelRow) GetQiHai() int32 { return pOwn.mQiHai } 
func (pOwn *sRofLevelRow) GetShenFa() int32 { return pOwn.mShenFa } 
func (pOwn *sRofLevelRow) GetSuDu() int32 { return pOwn.mSuDu } 
func (pOwn *sRofLevelRow) GetShangHaiJiaShen() float32 { return pOwn.mShangHaiJiaShen } 
func (pOwn *sRofLevelRow) GetShangHaiJianMian() float32 { return pOwn.mShangHaiJianMian } 
type sRofLevelTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofLevelRow
mRowMap map[int32]int32
}
func (pOwn *sRofLevelTable) newTypeObj() iRofRow {return new(sRofLevelRow)}
func (pOwn *sRofLevelTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofLevelTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofLevelTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofLevelRow)}
func (pOwn *sRofLevelTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofLevelTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofLevelRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofLevelTable) GetDataByID(aID int32) *sRofLevelRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofLevelTable) GetDataByRow(aIndex int32) *sRofLevelRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofLevelTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofLevelTable) GetCols() int32 {return pOwn.mColNum}
type sRofParameterRow struct {
mID int32
mParam1 int32
mParam2 string
}
func (pOwn *sRofParameterRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mParam1 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nParam2Len := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mParam2 = string(aBuffer[nOffset:nOffset+nParam2Len])
nOffset+=nParam2Len
return nOffset
}
func (pOwn *sRofParameterRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofParameterRow) GetParam1() int32 { return pOwn.mParam1 } 
func (pOwn *sRofParameterRow) GetParam2() string { return pOwn.mParam2 } 
type sRofParameterTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofParameterRow
mRowMap map[int32]int32
}
func (pOwn *sRofParameterTable) newTypeObj() iRofRow {return new(sRofParameterRow)}
func (pOwn *sRofParameterTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofParameterTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofParameterTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofParameterRow)}
func (pOwn *sRofParameterTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofParameterTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofParameterRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofParameterTable) GetDataByID(aID int32) *sRofParameterRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofParameterTable) GetDataByRow(aIndex int32) *sRofParameterRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofParameterTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofParameterTable) GetCols() int32 {return pOwn.mColNum}
type sRofSurnameRow struct {
mID int32
mFamilyName string
}
func (pOwn *sRofSurnameRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
nFamilyNameLen := int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mFamilyName = string(aBuffer[nOffset:nOffset+nFamilyNameLen])
nOffset+=nFamilyNameLen
return nOffset
}
func (pOwn *sRofSurnameRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofSurnameRow) GetFamilyName() string { return pOwn.mFamilyName } 
type sRofSurnameTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofSurnameRow
mRowMap map[int32]int32
}
func (pOwn *sRofSurnameTable) newTypeObj() iRofRow {return new(sRofSurnameRow)}
func (pOwn *sRofSurnameTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofSurnameTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofSurnameTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofSurnameRow)}
func (pOwn *sRofSurnameTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofSurnameTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofSurnameRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofSurnameTable) GetDataByID(aID int32) *sRofSurnameRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofSurnameTable) GetDataByRow(aIndex int32) *sRofSurnameRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofSurnameTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofSurnameTable) GetCols() int32 {return pOwn.mColNum}
type sRofThunderRow struct {
mID int32
mNum int32
mReduceHp int64
mFailLevel int32
mTime int32
mParam1 int32
mParam2 int32
}
func (pOwn *sRofThunderRow) readBody(aBuffer []byte) int32 {
var nOffset int32
pOwn.mID = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mNum = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mReduceHp = int64(binary.BigEndian.Uint64(aBuffer[nOffset:]))
nOffset+=8
pOwn.mFailLevel = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mTime = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mParam1 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
pOwn.mParam2 = int32(binary.BigEndian.Uint32(aBuffer[nOffset:]))
nOffset+=4
return nOffset
}
func (pOwn *sRofThunderRow) GetID() int32 { return pOwn.mID } 
func (pOwn *sRofThunderRow) GetNum() int32 { return pOwn.mNum } 
func (pOwn *sRofThunderRow) GetReduceHp() int64 { return pOwn.mReduceHp } 
func (pOwn *sRofThunderRow) GetFailLevel() int32 { return pOwn.mFailLevel } 
func (pOwn *sRofThunderRow) GetTime() int32 { return pOwn.mTime } 
func (pOwn *sRofThunderRow) GetParam1() int32 { return pOwn.mParam1 } 
func (pOwn *sRofThunderRow) GetParam2() int32 { return pOwn.mParam2 } 
type sRofThunderTable struct { 
mRowNum int32
mColNum int32
mIDMap  map[int32]*sRofThunderRow
mRowMap map[int32]int32
}
func (pOwn *sRofThunderTable) newTypeObj() iRofRow {return new(sRofThunderRow)}
func (pOwn *sRofThunderTable) setRowNum(aNum int32) {pOwn.mRowNum = aNum}
func (pOwn *sRofThunderTable) setColNum(aNum int32) {pOwn.mColNum = aNum}
func (pOwn *sRofThunderTable) setIDMap(aKey int32, aValue iRofRow) {pOwn.mIDMap[aKey] = aValue.(*sRofThunderRow)}
func (pOwn *sRofThunderTable) setRowMap(aKey int32, aValue int32) {pOwn.mRowMap[aKey] = aValue}
func (pOwn *sRofThunderTable) init(aPath string) bool {
pOwn.mIDMap = make(map[int32]*sRofThunderRow)
pOwn.mRowMap = make(map[int32]int32)
return analysisRof(aPath, pOwn)
}
func (pOwn *sRofThunderTable) GetDataByID(aID int32) *sRofThunderRow {return pOwn.mIDMap[aID]}
func (pOwn *sRofThunderTable) GetDataByRow(aIndex int32) *sRofThunderRow {
nID, ok := pOwn.mRowMap[aIndex]
if ok == false {return nil}
return pOwn.mIDMap[nID]
}
func (pOwn *sRofThunderTable) GetRows() int32 {return pOwn.mRowNum}
func (pOwn *sRofThunderTable) GetCols() int32 {return pOwn.mColNum}
