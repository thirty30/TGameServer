package logicserver

import (
	"runtime"
	"sync"
	"time"
)

type sMonitorElapsed struct {
	mFuncName    string
	mTotalTime   int64
	mLongestTime int64
	mTotalNum    int64
}

type sMonitor struct {
	mLock    sync.Mutex
	mTimeMap map[string]*sMonitorElapsed //耗时统计
}

func (pOwn *sMonitor) init() {
	pOwn.mTimeMap = make(map[string]*sMonitorElapsed)
	go pOwn.loop()
}

//defer recordTime(tag)()
func (pOwn *sMonitor) recordTime(aFuncName string) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start).Nanoseconds()
		pOwn.mLock.Lock()
		pElapsed := pOwn.mTimeMap[aFuncName]
		if pElapsed == nil {
			temp := new(sMonitorElapsed)
			temp.mFuncName = aFuncName
			temp.mTotalTime = duration
			temp.mTotalNum = 1
			pOwn.mTimeMap[aFuncName] = temp
		} else {
			pElapsed.mTotalTime += duration
			if duration > pElapsed.mLongestTime {
				pElapsed.mLongestTime = duration
			}
			pElapsed.mTotalNum++
		}
		pOwn.mLock.Unlock()
	}
}

func (pOwn *sMonitor) statisticTime() {
	pOwn.mLock.Lock()
	for _, v := range pOwn.mTimeMap {
		tTime := float32(v.mTotalTime) / float32(1000000)
		num := v.mTotalNum
		aveTime := tTime / float32(num)
		lTime := float32(v.mLongestTime) / float32(1000000)
		_LOG(LT_WARNING, "--Monitor-- elapsed time: %s total num: %d, total time:%.2fms, ave time: %.2fms, longest time: %.2fms", v.mFuncName, num, tTime, aveTime, lTime)
	}
	pOwn.mLock.Unlock()
}

func (pOwn *sMonitor) loop() {
	var memStats runtime.MemStats
	for {
		time.Sleep(time.Minute * 1)
		runtime.ReadMemStats(&memStats)
		pOwn.statisticTime()
		_LOG(LT_WARNING, "--Monitor-- current goroutine num:%d", runtime.NumGoroutine())
		memUse := float32(memStats.Alloc) / float32(1024) / float32(1024)
		_LOG(LT_WARNING, "--Monitor-- current use mem:%.2f MB", memUse)
	}
}
