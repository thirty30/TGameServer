package tdbmanager

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type sDBWorker struct {
	ID            int32
	mDBClient     *mongo.Client
	mDBSession    *mongo.Database
	mPreTaskList  chan *sDBTask
	mDoneTaskList chan *sDBTask
	mPlayerIDs    map[uint64]bool
}

func (pOwn *sDBWorker) init(aID int32, aDBSession *mongo.Database) {
	pOwn.ID = aID
	pOwn.mDBClient = aDBSession.Client()
	pOwn.mDBSession = aDBSession
	pOwn.mPreTaskList = make(chan *sDBTask, 10240)
	pOwn.mDoneTaskList = make(chan *sDBTask, 10240)
	pOwn.mPlayerIDs = make(map[uint64]bool)
}

func (pOwn *sDBWorker) start() {
	go pOwn.loop()
	go pOwn.ping()
}

func (pOwn *sDBWorker) clear() {
	pOwn.mDBClient.Disconnect(context.Background())
}

func (pOwn *sDBWorker) loop() {
	for {
		pTask := <-pOwn.mPreTaskList
		pCollection := pOwn.mDBSession.Collection(pTask.mCollectionName)

		switch pTask.mTaskType {
		case cDBTaskUpsert:
			{
				_, pTask.mError = pCollection.ReplaceOne(context.TODO(), pTask.mFilter, pTask.mData, options.Replace().SetUpsert(true))
			}
		case cDBTaskRemoveOne:
			{
				_, pTask.mError = pCollection.DeleteOne(context.TODO(), pTask.mFilter)
			}
		case cDBTaskRemove:
			{
				_, pTask.mError = pCollection.DeleteMany(context.TODO(), pTask.mFilter)
			}
		case cDBTaskFindOne:
			{
				pRes := pCollection.FindOne(context.TODO(), pTask.mFilter)
				pTask.mError = pRes.Decode(pTask.mData)
			}
		case cDBTaskFind:
			{
				pCursor, err := pCollection.Find(context.TODO(), pTask.mFilter)
				if err != nil {
					pTask.mError = err
					break
				}
				pTask.mError = pCursor.All(context.TODO(), pTask.mData)
			}
		case cDBTaskPing:
			{
				//pOwn.mDBClient.Ping(context.TODO(), nil)
			}
		case cDBTaskFinish:
			{

			}
		}
		if pTask.mCallBack != nil {
			pOwn.mDoneTaskList <- pTask
		}

	}
}

func (pOwn *sDBWorker) ping() {
	for {
		time.Sleep(time.Second * 60)
		pTask := new(sDBTask)
		pTask.mTaskType = cDBTaskPing
		pOwn.pushTask(pTask)
	}
}

func (pOwn *sDBWorker) pushTask(aTask *sDBTask) {
	pOwn.mPreTaskList <- aTask
}

func (pOwn *sDBWorker) addPlayer(aID uint64) {
	pOwn.mPlayerIDs[aID] = true
}

func (pOwn *sDBWorker) includePlayer(aID uint64) bool {
	_, bExist := pOwn.mPlayerIDs[aID]
	return bExist
}

func (pOwn *sDBWorker) removePlayer(aID uint64) {
	delete(pOwn.mPlayerIDs, aID)
}

func (pOwn *sDBWorker) GetActiveNum() int32 {
	return int32(len(pOwn.mPlayerIDs))
}
