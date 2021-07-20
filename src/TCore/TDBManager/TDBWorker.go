package tdbmanager

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	pOwn.mDBClient.Disconnect(ctx)
}

func (pOwn *sDBWorker) loop() {
	for {
		pTask := <-pOwn.mPreTaskList
		pCollection := pOwn.mDBSession.Collection(pTask.mCollectionName)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

		switch pTask.mTaskType {
		case cDBTaskUpsert:
			{
				//_, pTask.mError = pCollection.InsertOne(ctx, pTask.mData)
				filter := bson.M{"_id": pTask.mCondition}
				_, pTask.mError = pCollection.ReplaceOne(ctx, filter, pTask.mData, options.Replace().SetUpsert(true))
				if pTask.mError != nil {
					fmt.Println(pTask.mError.Error())
				}
			}
		case cDBTaskRemoveOne:
			{
				_, pTask.mError = pCollection.DeleteOne(ctx, pTask.mCondition)
			}
		case cDBTaskRemove:
			{
				_, pTask.mError = pCollection.DeleteMany(ctx, pTask.mCondition)
			}
		case cDBTaskFindOne:
			{
				pRes := pCollection.FindOne(ctx, pTask.mCondition)
				err := pRes.Decode(pTask.mData)
				if err != nil {
					pTask.mError = err
				}
			}
		case cDBTaskFind:
			{
				pCursor, err := pCollection.Find(ctx, pTask.mCondition)
				if err != nil {
					pTask.mError = err
					break
				}
				err = pCursor.All(ctx, pTask.mData)
				if err != nil {
					pTask.mError = err
				}
			}
		case cDBTaskPing:
			{
				//ctxPing, cancelPing := context.WithTimeout(context.Background(), time.Second*10)
				//cancelPing()
				//pOwn.mDBClient.Ping(ctxPing, nil)
			}
		}

		cancel()

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
