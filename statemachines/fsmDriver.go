package main

import (
	"fmt"

	"github.com/looplab/fsm"
	"github.com/sanshitsharma/golangsamples/statemachines/initfsm"
)

func main() {
	initFsm := initfsm.CreateInitFSM()

	visualization := fsm.Visualize(initFsm)
	fmt.Println(visualization)

	initFsm.Event(initfsm.StartM, "optional", "arguments")
	/*
		fmt.Println()
		initFsm.Event(initfsm.CritReady)
		fmt.Println()
		initFsm.Event(initfsm.CritFail)
		fmt.Println()
		initFsm.Event(initfsm.CritReady)
		fmt.Println()
		initFsm.Event(initfsm.OptReady)
	*/
}
