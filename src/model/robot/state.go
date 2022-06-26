package robot

type State struct {
	name   string // name of state
	intel  int    // how many abilities
	str    int    // melee, transport
	reflex int    // ranged
	speed  int    // how many squares per turn
	comm   int    // how far from HQ
}

func defaultRobotState() *State {
	return &State{
		name:  "robot",
		intel:  3,
		str:    3,
		reflex: 3,
		speed:  1,
		comm:   5,
	}
}
