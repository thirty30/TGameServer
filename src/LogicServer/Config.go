package logicserver

import "TCore"

type sConfig struct {
	ServerID  int32
	LogicPort int32
	DBPath    string
	DBPort    int32
	DBName    string
	DBUser    string
	DBPwd     string
	LogLevel  int32

	//debug mode
	GMEnalbe      int32
	IsDebug       int32
	MonitorEnable int32
}

//init config
func (pOwn *LogicServer) initConfig() bool {
	var config tcore.TConfig
	if config.Analysis("./Server.cfg") == false {
		return false
	}
	var isOK bool
	pOwn.mConfig.ServerID, isOK = config.GetDataInt32("ServerID")
	if isOK != true {
		return false
	}
	pOwn.mConfig.LogicPort, isOK = config.GetDataInt32("LogicPort")
	if isOK != true {
		return false
	}
	pOwn.mConfig.DBPath, isOK = config.GetDataString("DBPath")
	if isOK != true {
		return false
	}
	pOwn.mConfig.DBPort, isOK = config.GetDataInt32("DBPort")
	if isOK != true {
		return false
	}
	pOwn.mConfig.DBName, isOK = config.GetDataString("DBName")
	if isOK != true {
		return false
	}
	pOwn.mConfig.DBUser, isOK = config.GetDataString("DBUser")
	if isOK != true {
		return false
	}
	pOwn.mConfig.DBPwd, isOK = config.GetDataString("DBPwd")
	if isOK != true {
		return false
	}
	pOwn.mConfig.LogLevel, isOK = config.GetDataInt32("LogLevel")
	if isOK != true {
		return false
	}

	pOwn.mConfig.GMEnalbe, isOK = config.GetDataInt32("GMEnable")
	if isOK != true {
		return false
	}
	pOwn.mConfig.MonitorEnable, isOK = config.GetDataInt32("MonitorEnable")
	if isOK != true {
		return false
	}
	pOwn.mConfig.IsDebug, isOK = config.GetDataInt32("IsDebug")
	if isOK != true {
		return false
	}
	return true
}
