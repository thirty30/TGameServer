//Package tcore utility functions
package tcore

import (
	"math/rand"
	"time"
)

func GetNowTimeStamp() int64 {
	return time.Now().Unix()
}

//RandUint32 随机数
func RandUint32() uint32 {
	return rand.Uint32()
}

//RandUint32Range 随机范围
func RandUint32Range(aMin uint32, aMax uint32) uint32 {
	if aMax-aMin == 0 {
		return aMin
	}
	return rand.Uint32()%(aMax-aMin) + aMin
}

