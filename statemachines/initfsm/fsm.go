package initfsm

import (
	"fmt"

	"github.com/looplab/fsm"
)

// CreateInitFSM returns a new FSM object
func CreateInitFSM() *fsm.FSM {
	transitions := []fsm.EventDesc{
		{Name: Start, Src: []string{Idle}, Dst: Init},

		{Name: ModuleReady, Src: []string{Init}, Dst: Red},

		{Name: CritReady, Src: []string{Red}, Dst: Yellow},
		{Name: CritFail, Src: []string{Green, Yellow, Red}, Dst: Red},

		{Name: OptReady, Src: []string{Yellow}, Dst: Green},
		{Name: OptFail, Src: []string{Green, Yellow}, Dst: Yellow},
	}

	var initFsm = fsm.NewFSM(
		Idle,
		transitions,
		fsm.Callbacks{
			"before_event": func(e *fsm.Event) {
				fmt.Println("CB: Processing event:", e.Event, " src:", e.Src, " dst:", e.Dst, " count:", len(e.Args))
			},
			"after_event": func(e *fsm.Event) {
				fmt.Println("CB: Finished event:", e.Event)
			},
			"enter_state": func(e *fsm.Event) {
				fmt.Println("CB: entering state:", e.Dst)
			},
			"leave_state": func(e *fsm.Event) {
				fmt.Println("CB: Leaving state:", e.Src)
			},
		},
	)

	return initFsm
}
