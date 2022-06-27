package main

import (
	"github.com/honggim/robots/src/menu"
	"github.com/honggim/robots/src/model/robot"
	"github.com/honggim/robots/src/model/world"
	viewWorld "github.com/honggim/robots/src/view/world"
)

var planet *world.Planet
var bots []*robot.Robot

func init() {
	menu := menu.NewMenu()

	// login screen
	//menu.Login()

	// new / save / load game
	menu.File()

	// - map: size, features, i.e civ
	//m = models.GenerateMap(menu.GetMapOptions())
	planet = world.NewPlanet(menu.GetMockMapOptions())

	// - create core robots
	bots := menu.GetMockRobots()
	for i := 0; i < len(bots); i++ {
		bots[i].Render()
	}
}

func main() {
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
	//TODO:fmt.Scanln()
}
