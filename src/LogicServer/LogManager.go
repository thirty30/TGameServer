package logicserver

import (
	"TCore"
	"fmt"
	"os"
	"time"
)

type sLog struct {
	mFile     tcore.TFile
	mLogLevel int32
}

func (pOwn *sLog) init(aLogLevel int32) error {
	pOwn.mLogLevel = aLogLevel
	return pOwn.mFile.Init("./log/LogicServer", -1, tcore.TFileModeNew, tcore.TFileDay)
}

func (pOwn *sLog) formatTimeString() string {
	now := time.Now()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
}

func (pOwn *sLog) log(aFormat string, aParms ...interface{}) {
	str := pOwn.formatTimeString() + " [LOG] " + aFormat + "\n"
	if 1 >= pOwn.mLogLevel {
		pOwn.mFile.WriteFile(str, aParms...)
	}
	if gServerSingleton.getConfig().IsDebug == 1 {
		fmt.Fprintf(os.Stdout, str, aParms...)
	}
}

func (pOwn *sLog) warning(aFormat string, aParms ...interface{}) {
	str := pOwn.formatTimeString() + " [WARNING] " + aFormat + "\n"
	if 2 >= pOwn.mLogLevel {
		pOwn.mFile.WriteFile(str, aParms...)
	}
	if gServerSingleton.getConfig().IsDebug == 1 {
		fmt.Fprintf(os.Stdout, str, aParms...)
	}
}

func (pOwn *sLog) error(aFormat string, aParms ...interface{}) {
	str := pOwn.formatTimeString() + " [ERROR] " + aFormat + "\n"
	if 3 >= pOwn.mLogLevel {
		pOwn.mFile.WriteFile(str, aParms...)
	}
	if gServerSingleton.getConfig().IsDebug == 1 {
		fmt.Fprintf(os.Stdout, str, aParms...)
	}
}

func (pOwn *sLog) clear() {
	pOwn.mFile.Clear()
}
