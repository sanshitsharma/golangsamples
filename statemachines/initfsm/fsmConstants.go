package initfsm

const (
	// None state is an invalid state
	None string = "None"
	// Idle state is the start state, every new FSM instance is created in this state
	Idle string = "Idle"
	// Init state is entered upon receiving a start event
	Init string = "Init"
	// Red state is entered after module inits are complete
	Red string = "Red"
	// Yellow state is entered after module all critcal dependencies are determined live
	Yellow string = "Yellow"
	// Green state is entered after module all critcal & optional dependencies are determined live
	Green string = "Green"
)

// Events
const (
	// Start event kick starts the fsm
	Start string = "Start"
	// ModuleReady event is sent one all internal manager components are successfully initialized
	ModuleReady string = "ModuleReady"
	// CritReady event is sent once all the critical external dependencies' liveliness have been received
	CritReady string = "CritReady"
	// OptReady event is sent once all the optional external dependencies' liveliness have been received
	OptReady string = "OptReady"
	// CriticalFail event is sent when one of the critical dependencies' liveliness check fails
	CritFail string = "CritFail"
	// OptionalFail event is sent when one of the optional dependencies' liveliness check fails
	OptFail string = "OptFail"
)
