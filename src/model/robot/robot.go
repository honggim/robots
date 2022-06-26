package robot

import (
	"fmt"
)

type Robot struct {
	name     string
	states   []*State
	curState int
}

func NewRobot(name string) *Robot {
	state := defaultRobotState()
	robot := &Robot{
		name:      name,
		states:   []*State{state},
		curState: 0,
	}

	return robot
}

func (r *Robot) Name() string {
	return r.name
}

func (r *Robot) state() *State {
	return r.states[r.curState]
}

func (r *Robot) Intel() int {
	return r.state().intel
}

func (r *Robot) Speed() int {
	return r.state().speed
}

func (r *Robot) Reflexes() int {
	return r.state().reflex
}

func (r *Robot) Comm() int {
	return r.state().comm
}

func (r *Robot) Render() {
	fmt.Println(r.name)
	fmt.Println(r.state())
}

/*
func (r *Robot) States() {
}
*/

/*
func (r *Robot) Transform() {
}
*/

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
