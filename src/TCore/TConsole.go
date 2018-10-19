package tcore

import "fmt"

//CommandCallBack export
type CommandCallBack func(string)

//DebugConsole export
type DebugConsole struct {
	mCallBack CommandCallBack
}

//Init export
func (pOwn *DebugConsole) Init(aFunc CommandCallBack) {
	pOwn.mCallBack = aFunc
	go pOwn.loopInput()
}

func (pOwn *DebugConsole) loopInput() {
	for true {
		var cmd string
		fmt.Scanln(&cmd)
		pOwn.mCallBack(cmd)
	}
}
