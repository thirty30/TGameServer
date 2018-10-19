package tprotocol

import (
	"encoding/binary"
	"math"
)

//序列化
func serializeBOOL(aBuffer []byte, aSize uint32, aData bool) uint32 {
	if aSize < 1 {
		return 0
	}
	if aData == false {
		aBuffer[0] = byte(0)
	} else {
		aBuffer[0] = byte(1)
	}
	return 1
}

func serializeN8(aBuffer []byte, aSize uint32, aData int8) uint32 {
	if aSize < 1 {
		return 0
	}
	aBuffer[0] = byte(aData)
	return 1
}

func serializeN16(aBuffer []byte, aSize uint32, aData int16) uint32 {
	if aSize < 2 {
		return 0
	}
	binary.BigEndian.PutUint16(aBuffer, uint16(aData))
	return 2
}

func serializeN32(aBuffer []byte, aSize uint32, aData int32) uint32 {
	if aSize < 4 {
		return 0
	}
	binary.BigEndian.PutUint32(aBuffer, uint32(aData))
	return 4
}

func serializeN64(aBuffer []byte, aSize uint32, aData int64) uint32 {
	if aSize < 8 {
		return 0
	}
	binary.BigEndian.PutUint64(aBuffer, uint64(aData))
	return 8
}

func serializeU8(aBuffer []byte, aSize uint32, aData uint8) uint32 {
	if aSize < 1 {
		return 0
	}
	aBuffer[0] = byte(aData)
	return 1
}

func serializeU16(aBuffer []byte, aSize uint32, aData uint16) uint32 {
	if aSize < 2 {
		return 0
	}
	binary.BigEndian.PutUint16(aBuffer, aData)
	return 2
}

func serializeU32(aBuffer []byte, aSize uint32, aData uint32) uint32 {
	if aSize < 4 {
		return 0
	}
	binary.BigEndian.PutUint32(aBuffer, aData)
	return 4
}

func serializeU64(aBuffer []byte, aSize uint32, aData uint64) uint32 {
	if aSize < 8 {
		return 0
	}
	binary.BigEndian.PutUint64(aBuffer, aData)
	return 8
}

func serializeF32(aBuffer []byte, aSize uint32, aData float32) uint32 {
	if aSize < 4 {
		return 0
	}
	bits := math.Float32bits(aData)
	binary.BigEndian.PutUint32(aBuffer, bits)
	return 4
}

func serializeF64(aBuffer []byte, aSize uint32, aData float64) uint32 {
	if aSize < 8 {
		return 0
	}
	bits := math.Float64bits(aData)
	binary.BigEndian.PutUint64(aBuffer, bits)
	return 8
}

func serializeSTR(aBuffer []byte, aSize uint32, aData string) uint32 {
	nStrLen := len(aData)
	if aSize < uint32(4+nStrLen) {
		return 0
	}
	binary.BigEndian.PutUint32(aBuffer, uint32(nStrLen))
	copy(aBuffer[4:], []byte(aData))
	return uint32(4 + nStrLen)
}

//反序列化
func deserializeBOOL(aBuffer []byte, aSize uint32) (bool, uint32) {
	if aSize < 1 {
		return false, 0
	}
	res := int8(aBuffer[0])
	if res == 0 {
		return false, 1
	}
	return true, 1
}

func deserializeN8(aBuffer []byte, aSize uint32) (int8, uint32) {
	if aSize < 1 {
		return 0, 0
	}
	return int8(aBuffer[0]), 1
}

func deserializeN16(aBuffer []byte, aSize uint32) (int16, uint32) {
	if aSize < 2 {
		return 0, 0
	}
	return int16(binary.BigEndian.Uint16(aBuffer)), 2
}

func deserializeN32(aBuffer []byte, aSize uint32) (int32, uint32) {
	if aSize < 4 {
		return 0, 0
	}
	return int32(binary.BigEndian.Uint32(aBuffer)), 4
}

func deserializeN64(aBuffer []byte, aSize uint32) (int64, uint32) {
	if aSize < 8 {
		return 0, 0
	}
	return int64(binary.BigEndian.Uint64(aBuffer)), 8
}

func deserializeU8(aBuffer []byte, aSize uint32) (uint8, uint32) {
	if aSize < 1 {
		return 0, 0
	}
	return uint8(aBuffer[0]), 1
}

func deserializeU16(aBuffer []byte, aSize uint32) (uint16, uint32) {
	if aSize < 2 {
		return 0, 0
	}
	return binary.BigEndian.Uint16(aBuffer), 2
}

func deserializeU32(aBuffer []byte, aSize uint32) (uint32, uint32) {
	if aSize < 4 {
		return 0, 0
	}
	return binary.BigEndian.Uint32(aBuffer), 4
}

func deserializeU64(aBuffer []byte, aSize uint32) (uint64, uint32) {
	if aSize < 8 {
		return 0, 0
	}
	return binary.BigEndian.Uint64(aBuffer), 8
}

func deserializeF32(aBuffer []byte, aSize uint32) (float32, uint32) {
	if aSize < 4 {
		return 0, 0
	}
	bits := binary.BigEndian.Uint32(aBuffer)
	return math.Float32frombits(bits), 4
}

func deserializeF64(aBuffer []byte, aSize uint32) (float64, uint32) {
	if aSize < 8 {
		return 0, 0
	}
	bits := binary.BigEndian.Uint64(aBuffer)
	return math.Float64frombits(bits), 8
}

func deserializeSTR(aBuffer []byte, aSize uint32) (string, uint32) {
	if aSize < 4 {
		return "", 0
	}
	nStrLen := binary.BigEndian.Uint32(aBuffer)
	if aSize-4 < nStrLen {
		return "", 0
	}
	return string(aBuffer[4 : 4+nStrLen]), (nStrLen + 4)
}

//MessageBase exported
type MessageBase interface {
	Serialize(aBuffer []byte, aSize uint32) uint32
	Deserialize(aBuffer []byte, aSize uint32) uint32
}

//MessageHead exported
type MessageHead struct {
	MsgID        int32
	BodySize     uint32
	IsCompressed uint8
}

//GetHeadSize exported
func (pOwn *MessageHead) GetHeadSize() uint32 {
	return 9
}

//Serialize exported
func (pOwn *MessageHead) Serialize(aBuffer []byte, aSize uint32) uint32 {
	var nOffset uint32
	nOffset += serializeN32(aBuffer[nOffset:], aSize-nOffset, pOwn.MsgID)
	nOffset += serializeU32(aBuffer[nOffset:], aSize-nOffset, pOwn.BodySize)
	nOffset += serializeU8(aBuffer[nOffset:], aSize-nOffset, pOwn.IsCompressed)
	return nOffset
}

//Deserialize exported
func (pOwn *MessageHead) Deserialize(aBuffer []byte, aSize uint32) uint32 {
	var nOffset uint32
	var nTemp uint32
	pOwn.MsgID, nTemp = deserializeN32(aBuffer[nOffset:], aSize-nOffset)
	nOffset += nTemp
	pOwn.BodySize, nTemp = deserializeU32(aBuffer[nOffset:], aSize-nOffset)
	nOffset += nTemp
	pOwn.IsCompressed, nTemp = deserializeU8(aBuffer[nOffset:], aSize-nOffset)
	nOffset += nTemp
	return nOffset
}
