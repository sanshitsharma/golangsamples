package initfsm

// Init FSM States
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
	// Stop state is entered upon receiving a Abort event
	Stop string = "Stop"
)

// Init FSM Events
const (
	// Start event kick starts the fsm
	Start string = "Start"
	// ModuleReady event is sent one all internal manager components are successfully initialized
	ModuleReady string = "ModuleReady"
	// EssReady event is sent once all the critical dependencies' liveliness have been received
	EssReady string = "EssentialReady"
	// EssFail event is sent when one of the critical dependencies' liveliness check fails
	EssFail string = "EssentialFail"
	// AuxReady event is sent once all the secondary dependencies' liveliness have been received
	AuxReady string = "AuxiliaryReady"
	// AuxFail event is sent when one of the secondary dependencies' liveliness check fails
	AuxFail string = "AuxFail"
	// OptReady event is sent when all components in Green state are ready
	OptReady string = "OptionalReady"
	// OptFail event is sent when one or more green state components' liveliness check fails
	OptFail string = "OptionalFail"

	// Abort event is posted if a skeleton shutdown is received or basic init fails
	Abort string = "Abort"
)
