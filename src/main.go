package main

import (
	"fmt"
	"os"
	"time"

	gate "GateServer"
	logic "LogicServer"
)

type iServer interface {
	Init() bool
	Run()
	Clear()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("-[ logic | gate ]")
		return
	}

	switch os.Args[1] {
	case "-logic":
		pServer := new(logic.LogicServer)
		defer pServer.Clear()
		if pServer.Init() == true {
			pServer.Run()
		}
		break

	case "-gate":
		pServer := new(gate.GateServer)
		defer pServer.Clear()
		if pServer.Init() == true {
			pServer.Run()
		}
		break

	case "-debugall":
		startAll()
		break

	default:
		fmt.Println("wrong args!")
		return
	}

}

func startAll() {
	go func() {
		pServer := new(logic.LogicServer)
		if pServer.Init() == true {
			pServer.Run()
			pServer.Clear()
		}
	}()

	go func() {
		pServer := new(gate.GateServer)
		if pServer.Init() == true {
			pServer.Run()
			pServer.Clear()
		}
	}()

	for true {
		time.Sleep(time.Second)
	}
}
