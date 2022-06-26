package main

import (
	"github.com/honggim/robots/src/menu"
	"github.com/honggim/robots/src/model"
	"github.com/honggim/robots/src/model/robot"
	viewMap "github.com/honggim/robots/src/view/map"
)

var world *model.Map
var bots []*robot.Robot

func init() {
	menu := menu.NewMenu()

	// login screen
	//menu.Login()

	//TODO: menu for new game, save, load

	// - map: size, features, i.e civ
	//m = models.GenerateMap(menu.GetMapOptions())
	world = model.NewMap(menu.GetMockMapOptions())

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
	viewMap.Render(world)
	// - get resources
	// - per robot
	//   - move
	//   - ability
	//   - combat


	// base(s) actions

	// end turn?
	//TODO:fmt.Scanln()
}
