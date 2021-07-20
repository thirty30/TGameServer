package tdbmanager

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//TDBManager export
type TDBManager struct {
	mDBName              string
	mPlayerDataWorkerNum int32
	mPlayerDataWorkers   []*sDBWorker
	mPublicDataWorker    *sDBWorker
}

func (pOwn *TDBManager) SetWorkerNum(aNum int32) {
	pOwn.mPlayerDataWorkerNum = aNum
}

//Init export
func (pOwn *TDBManager) Init(aPath string, aPort int32, aAuth string, aDBName string, aDBUser string, aDBPwd string, aCryptoKey string) error {
	pOwn.mDBName = aDBName
	var err error
	realPwd, err := decryptPassword(aCryptoKey, aDBPwd)
	if err != nil {
		return err
	}

	//玩家数据worker
	if pOwn.mPlayerDataWorkerNum <= 0 {
		pOwn.mPlayerDataWorkerNum = 1
	}
	pOwn.mPlayerDataWorkers = make([]*sDBWorker, 0)
	for i := int32(1); i <= pOwn.mPlayerDataWorkerNum; i++ {
		clientOp := options.Client()
		clientOp.SetAuth(options.Credential{AuthSource: aAuth, Username: aDBUser, Password: realPwd})
		clientOp.ApplyURI("mongodb://" + aPath + ":" + strconv.Itoa(int(aPort)))
		client, err := mongo.NewClient(clientOp)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		err = client.Connect(ctx)
		if err != nil {
			cancel()
			return err
		}
		cancel()

		pWorker := new(sDBWorker)
		pOwn.mPlayerDataWorkers = append(pOwn.mPlayerDataWorkers, pWorker)

		pWorker.init(i, client.Database(aDBName))
		pWorker.start()
	}

	//公共数据worker
	{
		clientOp := options.Client()
		clientOp.SetAuth(options.Credential{AuthSource: aAuth, Username: aDBUser, Password: realPwd})
		clientOp.ApplyURI("mongodb://" + aPath + ":" + strconv.Itoa(int(aPort)))
		client, err := mongo.NewClient(clientOp)
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		err = client.Connect(ctx)
		if err != nil {
			cancel()
			return err
		}
		cancel()

		pOwn.mPublicDataWorker = new(sDBWorker)
		pOwn.mPublicDataWorker.init(-1, client.Database(aDBName))
		pOwn.mPublicDataWorker.start()
	}

	return nil
}

//Clear export
func (pOwn *TDBManager) Clear() {
	for _, v := range pOwn.mPlayerDataWorkers {
		v.clear()
	}
}

//EventDispatch export
func (pOwn *TDBManager) EventDispatch(aProcessNum int32) {
	nLimitCount := int32(0)
	for {
		bEmpty := true
		for _, v := range pOwn.mPlayerDataWorkers {
			select {
			case pTask := <-v.mDoneTaskList:
				{
					bEmpty = false
					if pTask.mCallBack == nil {
						break
					}
					pTask.mCallBack(pTask.mData, pTask.mCustomParm, pTask.mError)
					nLimitCount++
				}
			default:
				{
					break
				}
			}
		}

		if nLimitCount > aProcessNum || bEmpty == true {
			break
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Player data

func (pOwn *TDBManager) pushPlayerDataTask(
	aPlayerID uint64,
	aTaskType int8,
	aCollectionName string,
	aData interface{},
	aCallBack finishTaskCallBack,
	aCustomParm interface{}) {

	pTask := new(sDBTask)
	pTask.mTaskType = aTaskType
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aPlayerID
	pTask.mData = aData
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm

	for _, v := range pOwn.mPlayerDataWorkers {
		if v.includePlayer(aPlayerID) == true {
			v.pushTask(pTask)
			return
		}
	}

	//找一个人数最少的worker
	nNum := pOwn.mPlayerDataWorkers[0].GetActiveNum()
	pWorker := pOwn.mPlayerDataWorkers[0]
	for _, v := range pOwn.mPlayerDataWorkers {
		nCurNum := v.GetActiveNum()
		if nCurNum < nNum {
			nNum = nCurNum
			pWorker = v
		}
	}

	pWorker.addPlayer(aPlayerID)
	pWorker.pushTask(pTask)
}

//SavePlayerData export
func (pOwn *TDBManager) SavePlayerData(aPlayerID uint64, aCollectionName string, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushPlayerDataTask(aPlayerID, cDBTaskUpsert, aCollectionName, aData, aCallBack, aCustomParm)
}

//LoadPlayerData export
func (pOwn *TDBManager) LoadPlayerData(aPlayerID uint64, aCollectionName string, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushPlayerDataTask(aPlayerID, cDBTaskFindOne, aCollectionName, aData, aCallBack, aCustomParm)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//public data

func (pOwn *TDBManager) pushPublicDataTask(
	aTaskType int8,
	aCollectionName string,
	aCondition interface{},
	aData interface{},
	aCallBack finishTaskCallBack,
	aCustomParm interface{}) {

	pTask := new(sDBTask)
	pTask.mTaskType = aTaskType
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aCondition
	pTask.mData = aData
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm
	pOwn.mPublicDataWorker.pushTask(pTask)
}

//SavePublicDataByID export
func (pOwn *TDBManager) SavePublicDataByID(aCollectionName string, aID interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushPublicDataTask(cDBTaskUpsert, aCollectionName, aID, aData, aCallBack, aCustomParm)
}

//LoadPublicDataByID export
func (pOwn *TDBManager) LoadPublicDataByID(aCollectionName string, aID interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushPublicDataTask(cDBTaskFindOne, aCollectionName, aID, aData, aCallBack, aCustomParm)
}

//LoadPublicData export
func (pOwn *TDBManager) LoadPublicData(aCollectionName string, aCondition interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushPublicDataTask(cDBTaskFind, aCollectionName, aCondition, aData, aCallBack, aCustomParm)
}

//RemovePublicDataByID export
func (pOwn *TDBManager) RemovePublicDataByID(aCollectionName string, aID interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushPublicDataTask(cDBTaskRemoveOne, aCollectionName, aID, nil, aCallBack, aCustomParm)
}

//RemovePublicData export
func (pOwn *TDBManager) RemovePublicData(aCollectionName string, aCondition interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushPublicDataTask(cDBTaskRemove, aCollectionName, aCondition, nil, aCallBack, aCustomParm)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//保留接口
/*
func (pOwn *TDBManager) pushTask(
	aTaskType int8,
	aCollectionName string,
	aCondition interface{},
	aData interface{},
	aCallBack finishTaskCallBack,
	aCustomParm interface{}) {

	pTask := new(sDBTask)
	pTask.mTaskType = aTaskType
	pTask.mCollectionName = aCollectionName
	pTask.mCondition = aCondition
	pTask.mData = aData
	pTask.mCallBack = aCallBack
	pTask.mCustomParm = aCustomParm
	//pOwn.mPublicDataWorker.pushTask(pTask)
}

func (pOwn *TDBManager) upsertData(aCollectionName string, aID interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushTask(cDBTaskUpsert, aCollectionName, aID, aData, aCallBack, aCustomParm)
}


func (pOwn *TDBManager) findDataByID(aCollectionName string, aID interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushTask(cDBTaskFindOne, aCollectionName, aID, aData, aCallBack, aCustomParm)
}

func (pOwn *TDBManager) findData(aCollectionName string, aCondition interface{}, aData interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushTask(cDBTaskFind, aCollectionName, aCondition, aData, aCallBack, aCustomParm)
}

func (pOwn *TDBManager) removeDataByID(aCollectionName string, aID interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushTask(cDBTaskRemoveOne, aCollectionName, aID, nil, aCallBack, aCustomParm)
}

func (pOwn *TDBManager) removeData(aCollectionName string, aCondition interface{}, aCallBack finishTaskCallBack, aCustomParm interface{}) {
	pOwn.pushTask(cDBTaskRemove, aCollectionName, aCondition, nil, aCallBack, aCustomParm)
}

*/
