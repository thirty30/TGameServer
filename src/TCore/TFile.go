//Package tcore basic file
package tcore

import (
	"errors"
	"fmt"
	"os"
	"time"
)

//TFileCreateInterval 时间间隔
type TFileCreateInterval int64

//时间间隔单位
const (
	TFileSecond TFileCreateInterval = 1
	TFileMinute                     = TFileSecond * 60
	TFileHour                       = TFileMinute * 60
	TFileDay                        = TFileHour * 24
)

//文件打开方式
const (
	TFileModeNew    int32 = 1 //每次都新建
	TFileModeAppend int32 = 2 //追加
)

type sTFileBuffer struct {
	mBuffer []byte
}

func (pOwn *sTFileBuffer) Write(aBuf []byte) (int, error) {
	len := len(aBuf)
	pOwn.mBuffer = make([]byte, len)
	copy(pOwn.mBuffer, aBuf)
	return len, nil
}

//TFile export
type TFile struct {
	mChanWrite    chan *sTFileBuffer
	mChanNotify   chan int32
	mChanRecreate chan bool
	mChanHasClose chan bool
	mFileName     string
	mFileMode     int32
	mFileHandler  *os.File
	mInterval     TFileCreateInterval
}

//Init export
func (pOwn *TFile) Init(aFileName string, aPoolSize int32, aMode int32, aInterval TFileCreateInterval) error {
	pOwn.mFileName = aFileName
	pOwn.mFileMode = aMode
	var chanSize int32 = 1024
	if aPoolSize > 0 {
		chanSize = aPoolSize
	}
	pOwn.mChanWrite = make(chan *sTFileBuffer, chanSize)
	pOwn.mChanNotify = make(chan int32)
	pOwn.mChanRecreate = make(chan bool)
	pOwn.mChanHasClose = make(chan bool)

	pOwn.mInterval = aInterval
	if pOwn.mInterval > 0 {
		go pOwn.timer()
	}

	err := pOwn.openFile()
	if err != nil {
		return err
	}
	go pOwn.doWrite()
	return nil
}

func (pOwn *TFile) openFile() error {
	realFileName := pOwn.getRealFileName()
	var err error
	if pOwn.mFileMode == TFileModeNew {
		pOwn.mFileHandler, err = os.Create(realFileName)
		if err != nil {
			return err
		}
	} else if pOwn.mFileMode == TFileModeAppend {
		pOwn.mFileHandler, err = os.OpenFile(realFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
	} else {
		return errors.New("wrong file mode")
	}
	return nil
}

//WriteFile export
func (pOwn *TFile) WriteFile(aFormat string, aParms ...interface{}) error {
	pBuf := new(sTFileBuffer)
	_, err := fmt.Fprintf(pBuf, aFormat, aParms...)
	if err != nil {
		return err
	}
	pOwn.mChanWrite <- pBuf
	return nil
}

//Clear export
func (pOwn *TFile) Clear() {
	pOwn.mChanNotify <- 2
}

//write data to real file
func (pOwn *TFile) doWrite() {
	for {
		select {
		case pTempBuf := <-pOwn.mChanWrite:
			pOwn.mFileHandler.Write(pTempBuf.mBuffer)

		case op := <-pOwn.mChanNotify:
			if op == 1 {
				pOwn.mFileHandler.Close()
				pOwn.openFile()
				pOwn.mChanRecreate <- true
			} else if op == 2 {
				pOwn.mFileHandler.Close()
				pOwn.mChanHasClose <- true //todo: close the timer goroutine
			}

		}
	}
}

func (pOwn *TFile) getRealFileName() string {
	if pOwn.mInterval > 0 {
		pBuf := new(sTFileBuffer)
		now := time.Now()
		fmt.Fprintf(pBuf, "%s%d_%d_%d_%d_%d_%d", pOwn.mFileName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
		return string(pBuf.mBuffer)
	}
	return pOwn.mFileName
}

func (pOwn *TFile) timer() {
	lastTime := time.Now().Unix()
	for {
		if time.Now().Unix() > (lastTime + int64(pOwn.mInterval)) {
			pOwn.mChanNotify <- 1
			<-pOwn.mChanRecreate
			lastTime = time.Now().Unix()
		}
		time.Sleep(time.Second)
	}
}
