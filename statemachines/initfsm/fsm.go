package initfsm

import (
	"fmt"

	"github.com/looplab/fsm"
)

func beforeStartCb(e *fsm.Event) {
	fmt.Println("beforeStart CB. event:", e.Event, " src:", e.Src, " dst:", e.Dst, " args:", e.Args)

	if opt, ok := e.Args[0].(string); !ok {
		fmt.Println("invalid argument type")
	} else {
		fmt.Printf("Argument '%v' is a string\n", opt)
	}
}

// CreateInitFSM returns a new FSM object
func CreateInitFSM() *fsm.FSM {
	transitions := []fsm.EventDesc{
		{Name: StartM, Src: []string{Idle}, Dst: Red},

		{Name: VitalReady, Src: []string{Red, Yellow}, Dst: Yellow},
		{Name: VitalFail, Src: []string{Red, Yellow, Green}, Dst: Red},

		{Name: EssReady, Src: []string{Yellow, Green}, Dst: Green},
		{Name: EssFail, Src: []string{Yellow, Green}, Dst: Yellow},

		{Name: Shutdown, Src: []string{Idle, Red, Yellow, Green}, Dst: Stop},
	}
	var initFsm = fsm.NewFSM(
		Idle,
		transitions,
		fsm.Callbacks{
			// Common callbacks
			"before_event": func(e *fsm.Event) {
				fmt.Println("CB: Processing event. event:", e.Event, "src:", e.Src, "dst:", e.Dst, "args:", e.Args)
			},
			"after_event": func(e *fsm.Event) {
				fmt.Println("CB: Finished event. event:", e.Event)
			},
			"enter_state": func(e *fsm.Event) {
				fmt.Println("CB: Entering state. state:", e.Dst)
			},
			"leave_state": func(e *fsm.Event) {
				fmt.Println("CB: Leaving state. state:", e.Src)
			},

			// before_<event> callbacks
			"before_Start": beforeStartCb,
		},
	)
	
	return initFsm
}
