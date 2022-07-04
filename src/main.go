package main

import (
	"fmt"

	"github.com/honggim/robots/src/menu"
	//"github.com/honggim/robots/src/model/robot"
	//"github.com/honggim/robots/src/model/world"
	viewWorld "github.com/honggim/robots/src/view/world"
)

func main() {
	m := menu.NewMenu()

	// login screen
	//m.Login()

	// new / save / load game
	// -> new game
	// - map: size, features, i.e civ
	//planet = world.NewPlanet(m.GetMockMapOptions())

	// -> load game
	var filename string = "savefile.txt"
	planet := m.Load(filename)

	// - create core robots
	/*
	var bots []*robot.Robot
	bots := m.GetMockRobots()
	for i := 0; i < len(bots); i++ {
		bots[i].Render()
	}
	*/

	//TODO: move into controller

	// per user turn
	// - render map
	viewWorld.Render(planet)
	// - get resources
	// - per robot
	//   - move
	//   - ability
	//   - combat


	// base(s) actions

	// end turn?

	// -> save game
	//m.Save(filename, planet)
	fmt.Scanln()
}
