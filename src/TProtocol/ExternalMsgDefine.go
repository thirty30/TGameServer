package tprotocol
type CommonBool struct{
Value bool
}
func (pOwn *CommonBool) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeBOOL(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonBool) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeBOOL(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN8 struct{
Value int8
}
func (pOwn *CommonN8) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN8) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN16 struct{
Value int16
}
func (pOwn *CommonN16) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN16) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN32 struct{
Value int32
}
func (pOwn *CommonN32) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN32) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonN64 struct{
Value int64
}
func (pOwn *CommonN64) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonN64) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeN64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU8 struct{
Value uint8
}
func (pOwn *CommonU8) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU8) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU16 struct{
Value uint16
}
func (pOwn *CommonU16) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU16) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU32 struct{
Value uint32
}
func (pOwn *CommonU32) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU32) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonU64 struct{
Value uint64
}
func (pOwn *CommonU64) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonU64) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonF32 struct{
Value float32
}
func (pOwn *CommonF32) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonF32) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonF64 struct{
Value float64
}
func (pOwn *CommonF64) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeF64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonF64) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeF64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type CommonStr struct{
Value string
}
func (pOwn *CommonStr) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value)
return nOffset
}
func (pOwn *CommonStr) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type Login struct{
UID string
}
func (pOwn *Login) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.UID)
return nOffset
}
func (pOwn *Login) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.UID, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type LoginResp struct{
ResCode uint8
PlayerID uint64
Renamed bool
PlayerName string
Level int32
Exp uint64
QiYun int32
Age uint64
SelectFaction int32
HierarchyLv int32
}
func (pOwn *LoginResp) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU8(aBuffer[nOffset:], aSize-nOffset, pOwn.ResCode)
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.PlayerID)
nOffset += serializeBOOL(aBuffer[nOffset:], aSize-nOffset, pOwn.Renamed)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.PlayerName)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Level)
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.Exp)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.QiYun)
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.Age)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.SelectFaction)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.HierarchyLv)
return nOffset
}
func (pOwn *LoginResp) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.ResCode, nTemp = deserializeU8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.PlayerID, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Renamed, nTemp = deserializeBOOL(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.PlayerName, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Level, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Exp, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.QiYun, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Age, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.SelectFaction, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.HierarchyLv, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type ChangeName struct{
PlayerName string
}
func (pOwn *ChangeName) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.PlayerName)
return nOffset
}
func (pOwn *ChangeName) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.PlayerName, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type SyncLevel struct{
Level int32
Exp uint64
QiYun int32
}
func (pOwn *SyncLevel) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Level)
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.Exp)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.QiYun)
return nOffset
}
func (pOwn *SyncLevel) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Level, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Exp, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.QiYun, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type SyncXiuLianInfo struct{
TotalLingQi int32
ItemID int32
GetTime int64
EndTime int64
}
func (pOwn *SyncXiuLianInfo) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.TotalLingQi)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.ItemID)
nOffset += serializeN64(aBuffer[nOffset:], aSize-nOffset, pOwn.GetTime)
nOffset += serializeN64(aBuffer[nOffset:], aSize-nOffset, pOwn.EndTime)
return nOffset
}
func (pOwn *SyncXiuLianInfo) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.TotalLingQi, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.ItemID, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.GetTime, nTemp = deserializeN64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.EndTime, nTemp = deserializeN64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type HierarchyUp struct{
HierarchyID int32
}
func (pOwn *HierarchyUp) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.HierarchyID)
return nOffset
}
func (pOwn *HierarchyUp) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.HierarchyID, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type SelectFaction struct{
FactionID int32
}
func (pOwn *SelectFaction) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.FactionID)
return nOffset
}
func (pOwn *SelectFaction) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.FactionID, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type FightInfo struct{
RedPlayerID uint64
BluePlayerID uint64
RedPlayerName string
BluePlayerName string
Win int32
}
func (pOwn *FightInfo) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.RedPlayerID)
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.BluePlayerID)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.RedPlayerName)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.BluePlayerName)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Win)
return nOffset
}
func (pOwn *FightInfo) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.RedPlayerID, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.BluePlayerID, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.RedPlayerName, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.BluePlayerName, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Win, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type OnlinePlayerInfo struct{
PlayerID uint64
PlayerName string
}
func (pOwn *OnlinePlayerInfo) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.PlayerID)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.PlayerName)
return nOffset
}
func (pOwn *OnlinePlayerInfo) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.PlayerID, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.PlayerName, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type AllOnlinePlayerInfo struct{
mAllPlayer []*OnlinePlayerInfo
}
func (pOwn *AllOnlinePlayerInfo) GetAllPlayerCount() int32 { return int32(len(pOwn.mAllPlayer)) }
func (pOwn *AllOnlinePlayerInfo) GetAllPlayerAt(aIdx int32) *OnlinePlayerInfo { return pOwn.mAllPlayer[aIdx] }
func (pOwn *AllOnlinePlayerInfo) AppendAllPlayer(aData *OnlinePlayerInfo) {
 if pOwn.mAllPlayer == nil { pOwn.mAllPlayer = make([]*OnlinePlayerInfo, 0, 8) }
 pOwn.mAllPlayer = append(pOwn.mAllPlayer, aData)
 }
func (pOwn *AllOnlinePlayerInfo) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nAllPlayerCount := len(pOwn.mAllPlayer)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(nAllPlayerCount))
for i := 0; i < nAllPlayerCount; i++ { nOffset += pOwn.mAllPlayer[i].Serialize(aBuffer[nOffset:], aSize-nOffset) }
return nOffset
}
func (pOwn *AllOnlinePlayerInfo) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
nAllPlayerCount, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
for i := 0; i < int(nAllPlayerCount); i++ {
var obj OnlinePlayerInfo
nOffset += obj.Deserialize(aBuffer[nOffset:], aSize-nOffset)
if pOwn.mAllPlayer == nil { pOwn.mAllPlayer = make([]*OnlinePlayerInfo, 0, 8) }
pOwn.mAllPlayer = append(pOwn.mAllPlayer, &obj)
}
return nOffset
}
type T1 struct{
Value1 int32
Value2 float32
Value3 string
}
func (pOwn *T1) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value1)
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value2)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value3)
return nOffset
}
func (pOwn *T1) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value1, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value2, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value3, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type T2 struct{
Value1 int32
Value2 float32
Value3 string
Value4 int32
}
func (pOwn *T2) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value1)
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value2)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value3)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value4)
return nOffset
}
func (pOwn *T2) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value1, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value2, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value3, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value4, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
return nOffset
}
type Test struct{
Value0 bool
Value1 int8
Value2 int16
Value3 int32
Value4 int64
Value5 uint8
Value6 uint16
Value7 uint32
Value8 uint64
Value9 float32
Value10 float64
Value11 string
Value12 T1
mValue13 []*T2
mValue14 []int32
mValue15 []string
}
func (pOwn *Test) GetValue13Count() int32 { return int32(len(pOwn.mValue13)) }
func (pOwn *Test) GetValue13At(aIdx int32) *T2 { return pOwn.mValue13[aIdx] }
func (pOwn *Test) AppendValue13(aData *T2) {
 if pOwn.mValue13 == nil { pOwn.mValue13 = make([]*T2, 0, 8) }
 pOwn.mValue13 = append(pOwn.mValue13, aData)
 }
func (pOwn *Test) GetValue14Count() int32 { return int32(len(pOwn.mValue14)) }
func (pOwn *Test) GetValue14At(aIdx int32) int32 { return pOwn.mValue14[aIdx] }
func (pOwn *Test) AppendValue14(aData int32) {
 if pOwn.mValue14 == nil { pOwn.mValue14 = make([]int32, 0, 8) }
 pOwn.mValue14 = append(pOwn.mValue14, aData)
 }
func (pOwn *Test) GetValue15Count() int32 { return int32(len(pOwn.mValue15)) }
func (pOwn *Test) GetValue15At(aIdx int32) string { return pOwn.mValue15[aIdx] }
func (pOwn *Test) AppendValue15(aData string) {
 if pOwn.mValue15 == nil { pOwn.mValue15 = make([]string, 0, 8) }
 pOwn.mValue15 = append(pOwn.mValue15, aData)
 }
func (pOwn *Test) Serialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
nOffset += serializeBOOL(aBuffer[nOffset:], aSize-nOffset, pOwn.Value0)
nOffset += serializeN8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value1)
nOffset += serializeN16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value2)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value3)
nOffset += serializeN64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value4)
nOffset += serializeU8(aBuffer[nOffset:], aSize-nOffset, pOwn.Value5)
nOffset += serializeU16(aBuffer[nOffset:], aSize-nOffset, pOwn.Value6)
nOffset += serializeU32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value7)
nOffset += serializeU64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value8)
nOffset += serializeF32(aBuffer[nOffset:], aSize-nOffset, pOwn.Value9)
nOffset += serializeF64(aBuffer[nOffset:], aSize-nOffset, pOwn.Value10)
nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.Value11)
nOffset += pOwn.Value12.Serialize(aBuffer[nOffset:], aSize-nOffset)
nValue13Count := len(pOwn.mValue13)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(nValue13Count))
for i := 0; i < nValue13Count; i++ { nOffset += pOwn.mValue13[i].Serialize(aBuffer[nOffset:], aSize-nOffset) }
nValue14Count := len(pOwn.mValue14)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(nValue14Count))
for i := 0; i < nValue14Count; i++ { nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.mValue14[i]) }
nValue15Count := len(pOwn.mValue15)
nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, int32(nValue15Count))
for i := 0; i < nValue15Count; i++ { nOffset += serializeSTR(aBuffer[nOffset:], aSize-nOffset, pOwn.mValue15[i]) }
return nOffset
}
func (pOwn *Test) Deserialize(aBuffer []byte, aSize uint32) uint32{
var nOffset uint32
var nTemp uint32
pOwn.Value0, nTemp = deserializeBOOL(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value1, nTemp = deserializeN8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value2, nTemp = deserializeN16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value3, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value4, nTemp = deserializeN64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value5, nTemp = deserializeU8(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value6, nTemp = deserializeU16(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value7, nTemp = deserializeU32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value8, nTemp = deserializeU64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value9, nTemp = deserializeF32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value10, nTemp = deserializeF64(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
pOwn.Value11, nTemp = deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
nOffset += pOwn.Value12.Deserialize(aBuffer[nOffset:], aSize-nOffset)
nValue13Count, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
for i := 0; i < int(nValue13Count); i++ {
var obj T2
nOffset += obj.Deserialize(aBuffer[nOffset:], aSize-nOffset)
if pOwn.mValue13 == nil { pOwn.mValue13 = make([]*T2, 0, 8) }
pOwn.mValue13 = append(pOwn.mValue13, &obj)
}
nValue14Count, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
for i := 0; i < int(nValue14Count); i++ {
obj, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
if pOwn.mValue14 == nil { pOwn.mValue14 = make([]int32, 0, 8) }
pOwn.mValue14 = append(pOwn.mValue14, obj)
}
nValue15Count, nTemp := deserializeN32(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
for i := 0; i < int(nValue15Count); i++ {
obj, nTemp := deserializeSTR(aBuffer[nOffset:], aSize-nOffset)
nOffset += nTemp
if pOwn.mValue15 == nil { pOwn.mValue15 = make([]string, 0, 8) }
pOwn.mValue15 = append(pOwn.mValue15, obj)
}
return nOffset
}
const (
C2S_LOGIN = 1
S2C_LOGIN_RESP = 2
C2S_CHANGE_NAME = 3
S2C_CHANGE_NAME_RESP = 4
C2S_XIULIAN_INFO = 5
C2S_START_XIULIAN = 6
S2C_SYNC_XIULIAN_INFO = 7
C2S_GET_XIULIAN_EXP = 8
C2S_LEVEL_UP = 9
S2C_SYNC_LEVEL = 10
S2C_SYNC_AGE = 11
C2S_GET_SERVER_TIME = 12
S2C_SYNC_SERVER_TIME = 13
C2S_HIERARCHY_UP = 14
S2C_HIERARCHY_UP_RESP = 15
C2S_SELECTFACTION = 16
S2C_SELECTFACTION_RESP = 17
C2S_ATK_PLAYER = 18
S2C_ATK_PLAYER_RESP = 19
C2S_GET_ALL_ONLINE_PLAYER = 20
S2C_GET_ALL_ONLINE_PLAYER_RESP = 21
C2S_TEST = 22
S2C_TEST_RESP = 23
)
