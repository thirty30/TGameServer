package tdbmanager

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

const (
	cDBTaskUpsert = iota
	cDBTaskFindOne
	cDBTaskFind
	cDBTaskRemoveOne
	cDBTaskRemove
	cDBTaskPing
	cDBTaskFinish
)

type finishTaskCallBack func(aResData interface{}, aCustomParm interface{}, aError error)

type sDBTask struct {
	mTaskType       int8
	mCollectionName string
	mFilter         interface{}
	mData           interface{}
	mCallBack       finishTaskCallBack
	mCustomParm     interface{}
	mError          error
}

//key 和 password 都必须是16位
func encryptPassword(aKey string, aPassword string) (string, error) {
	block, err := aes.NewCipher([]byte(aKey))
	if err != nil {
		return "", err
	}
	iv := make([]byte, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	src := []byte(aPassword)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return hex.EncodeToString(dst), nil
}

func decryptPassword(aKey string, aPassword string) (string, error) {
	src, err := hex.DecodeString(aPassword)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(aKey))
	if err != nil {
		return "", err
	}
	iv := make([]byte, block.BlockSize())
	blockMode := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return string(dst), nil
}
