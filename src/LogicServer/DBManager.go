package logicserver

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
)

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

const (
	cDBTaskUpsert = iota
	cDBTaskFindOne
	cDBTaskFind
	cDBTaskRemoveOne
	cDBTaskRemove
)

type finishTaskCallBack func(aResData interface{}, aCustomParm interface{}, aError error)

type sDBTask struct {
	mTaskType       int8
	mCollectionName string
	mCondition      interface{}
	mData           interface{}
	mCallBack       finishTaskCallBack
	mCustomParm     interface{}
	mError          error
}

//DBManager export
type sDBManager struct {
	mDBSession    *mgo.Session
	mDBName       string
	mPreTaskList  chan *sDBTask
	mDoneTaskList chan *sDBTask
}

//Init export
func (pOwn *sDBManager) init(aPath string, aPort int32, aDBName string, aDBUser string, aDBPwd string, aCryptoKey string) error {
	var err error
	realPwd, err := decryptPassword(aCryptoKey, aDBPwd)
	if err != nil {
		return err
	}
	dialFnfo := &mgo.DialInfo{
		Addrs:     []string{aPath + ":" + strconv.Itoa(int(aPort))},
		Direct:    false,
		Timeout:   time.Second * 5,
		Database:  aDBName,
		Source:    "admin",
		Username:  aDBUser,
		Password:  realPwd,
		PoolLimit: 2048,
	}
	pOwn.mDBSession, err = mgo.DialWithInfo(dialFnfo)
	if err != nil {
		return err
	}
	pOwn.mDBSession.SetMode(mgo.Monotonic, true)
	pOwn.mDBName = aDBName
	pOwn.mPreTaskList = make(chan *sDBTask, 10240)
	pOwn.mDoneTaskList = make(chan *sDBTask, 10240)
	go pOwn.loop()
	return nil
}

//Clear export
func (pOwn *sDBManager) clear() {
	if pOwn.mDBSession != nil {
		pOwn.mDBSession.Close()
	}
}

//EventDispatch export
func (pOwn *sDBManager) eventDispatch() {
	nLimitCount := 0
	select {
	case pTask := <-pOwn.mDoneTaskList:
		if pTask.mCallBack != nil {
			pTask.mCallBack(pTask.mData, pTask.mCustomParm, pTask.mError)
		}
		nLimitCount++
		if nLimitCount > 100 {
			break
		}
	default:
		break
	}
}

func (pOwn *sDBManager) loop() {
	for true {
		pTask := <-pOwn.mPreTaskList
		// pTempSession := pOwn.mDBSession.Copy()
		// defer pTempSession.Close()
		pCollection := pOwn.mDBSession.DB(pOwn.mDBName).C(pTask.mCollectionName)

		switch pTask.mTaskType {
		case cDBTaskUpsert:
			{
				_, pTask.mError = pCollection.UpsertId(pTask.mCondition, pTask.mData)
				break
			}
		case cDBTaskRemoveOne:
			{
				pTask.mError = pCollection.RemoveId(pTask.mCondition)
				break
			}
		case cDBTaskRemove:
			{
				pTask.mError = pCollection.Remove(pTask.mCondition)
				break
			}
		case cDBTaskFindOne:
			{
				pQuery := pCollection.FindId(pTask.mCondition)
				err := pQuery.One(pTask.mData)
				if err != nil && err != mgo.ErrNotFound {
					pTask.mError = err
				}
				break
			}
		case cDBTaskFind:
			{
				pQuery := pCollection.Find(pTask.mCondition)
				err := pQuery.All(pTask.mData)
				if err != nil && err != mgo.ErrNotFound {
					pTask.mError = err
				}
				break
			}
		}

		if pTask.mCallBack != nil {
			pOwn.mDoneTaskList <- pTask
		}

		pOwn.mDBSession.Ping()
	}
}

func (pOwn *sDBManager) pushTask(aTask *sDBTask) {
	pOwn.mPreTaskList <- aTask
}

//UpsertData export
func (pOwn *sDBManager) upsertData(aCollectionName string, aID interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pTask := new(sDBTask)
	pTask.mTaskType = cDBTaskUpsert
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aID
	pTask.mData = aData
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm
	pOwn.pushTask(pTask)
}

//RemoveDataByID export
func (pOwn *sDBManager) removeDataByID(aCollectionName string, aID interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pTask := new(sDBTask)
	pTask.mTaskType = cDBTaskRemoveOne
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aID
	pTask.mData = nil
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm
	pOwn.pushTask(pTask)
}

//RemoveData export
func (pOwn *sDBManager) removeData(aCollectionName string, aCondition interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pTask := new(sDBTask)
	pTask.mTaskType = cDBTaskRemove
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aCondition
	pTask.mData = nil
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm
	pOwn.pushTask(pTask)
}

//FindDataByID export
func (pOwn *sDBManager) findDataByID(aCollectionName string, aID interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pTask := new(sDBTask)
	pTask.mTaskType = cDBTaskFindOne
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aID
	pTask.mData = aData
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm
	pOwn.pushTask(pTask)
}

//FindData export
func (pOwn *sDBManager) findData(aCollectionName string, aCondition interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pTask := new(sDBTask)
	pTask.mTaskType = cDBTaskFind
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aCondition
	pTask.mData = aData
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm
	pOwn.pushTask(pTask)
}
