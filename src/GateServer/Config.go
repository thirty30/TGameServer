package gateserver

import "TCore"

type sConfig struct {
	ExternalPort int32
	InternalPort int32
	LogicPort    int32
	LogLevel     int32

	IsDebug int32
}

//init config
func (pOwn *GateServer) initConfig() bool {
	var config tcore.TConfig
	if config.Analysis("./Server.cfg") == false {
		return false
	}
	var isOK bool
	pOwn.mConfig.ExternalPort, isOK = config.GetDataInt32("ExternalGatePort")
	if isOK != true {
		return false
	}
	pOwn.mConfig.InternalPort, isOK = config.GetDataInt32("InternalGatePort")
	if isOK != true {
		return false
	}
	pOwn.mConfig.LogicPort, isOK = config.GetDataInt32("LogicPort")
	if isOK != true {
		return false
	}
	pOwn.mConfig.LogLevel, isOK = config.GetDataInt32("LogLevel")
	if isOK != true {
		return false
	}
	pOwn.mConfig.IsDebug, isOK = config.GetDataInt32("IsDebug")
	if isOK != true {
		return false
	}
	return true
}
