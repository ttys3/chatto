package fsm

// Config models the yaml configuration
type Config struct {
	States    []string          `yaml:"states"`
	Commands  []string          `yaml:"commands"`
	Functions []Function        `yaml:"functions"`
	Defaults  map[string]string `yaml:"defaults"`
}

// Function models a function in yaml
type Function struct {
	Tuple      Tuple  `yaml:"tuple"`
	Transition string `yaml:"transition"`
	Message    string `yaml:"message"`
}

// Tuple models a command state tuple in yaml
type Tuple struct {
	Command string `yaml:"command"`
	State   string `yaml:"state"`
}

// Domain models the final configuration of an FSM
type Domain struct {
	StateTable      map[string]int
	CommandList     []string
	TransitionTable map[CmdStateTupple]TransitionFunc
	DefaultMessages map[string]string
}

// CmdStateTupple is a tuple of Command and State
type CmdStateTupple struct {
	Cmd   string
	State int
}

// TransitionFunc models a transition function
type TransitionFunc Transition

// Transition models a state transition
type Transition struct {
	Next    int
	Message string
}

// FSM models a Finite State Machine
type FSM struct {
	State int
}

// ExecuteCmd executes a command in FSM
func (m *FSM) ExecuteCmd(cmd string, dom Domain) string {
	tupple := CmdStateTupple{cmd, m.State}
	trans := dom.TransitionTable[tupple]
	if trans == (TransitionFunc{}) {
		return dom.DefaultMessages["unknown"]
	}
	m.State = trans.Next
	return trans.Message
}