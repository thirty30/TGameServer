//Package tcore log file
package tcore

import (
	"fmt"
	"os"
	"time"
)

type TLog struct {
	mFilePath string
	mLogLevel int32
	mIsPrint  bool
	mFile     TFile
}

func (pOwn *TLog) Init(aFilePath string, aLogLevel int32, aIsPrint bool) error {
	pOwn.mFilePath = aFilePath
	pOwn.mLogLevel = aLogLevel
	pOwn.mIsPrint = aIsPrint

	return pOwn.mFile.Init(pOwn.mFilePath, -1, TFileModeNew, TFileDay)
}

func (pOwn *TLog) formatTimeString() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func (pOwn *TLog) Log(aFormat string, aParms ...interface{}) {
	str := pOwn.formatTimeString() + " [LOG] " + aFormat + "\n"
	if 1 >= pOwn.mLogLevel {
		pOwn.mFile.WriteFile(str, aParms...)
	}
	if pOwn.mIsPrint == true {
		fmt.Fprintf(os.Stdout, str, aParms...)
	}
}

func (pOwn *TLog) Warning(aFormat string, aParms ...interface{}) {
	str := pOwn.formatTimeString() + " [WARNING] " + aFormat + "\n"
	if 2 >= pOwn.mLogLevel {
		pOwn.mFile.WriteFile(str, aParms...)
	}
	if pOwn.mIsPrint == true {
		fmt.Fprintf(os.Stdout, str, aParms...)
	}
}

func (pOwn *TLog) Error(aFormat string, aParms ...interface{}) {
	str := pOwn.formatTimeString() + " [ERROR] " + aFormat + "\n"
	if 3 >= pOwn.mLogLevel {
		pOwn.mFile.WriteFile(str, aParms...)
	}
	if pOwn.mIsPrint == true {
		fmt.Fprintf(os.Stdout, str, aParms...)
	}
}

func (pOwn *TLog) Clear() {
	pOwn.mFile.Clear()
}
