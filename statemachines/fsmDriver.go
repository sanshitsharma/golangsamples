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

	initFsm.Event(initfsm.Start)
}
