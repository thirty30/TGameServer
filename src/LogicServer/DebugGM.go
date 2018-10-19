package logicserver

import (
	"strings"
)

type sDebugGM struct {
}

func (pOwn *sDebugGM) dispatchCommand(aPlayerID string, aCommand string) {
	parm := strings.Split(aCommand, "=")
	switch parm[0] {
	case "addItem":
		{
			pOwn.addItem(aPlayerID, parm[1:])
		}
		break
	}
}

func (pOwn *sDebugGM) addItem(aPlayerID string, aParm []string) {

}
