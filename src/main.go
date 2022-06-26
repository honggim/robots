package main

import (
	"github.com/honggim/robots/src/menu"
	"github.com/honggim/robots/src/model"
	viewMap "github.com/honggim/robots/src/view/map"
)

var m *model.Map

func init() {
	menu := menu.NewMenu()

	// login screen
	//menu.Login()

	//TODO: menu for new game, save, load

	// - map: size, features, i.e civ
	//m = models.GenerateMap(menu.GetMapOptions())
	m = model.NewMap(menu.GetMockMapOptions())

	// - pick core robots
}

func main() {
	//TODO: move into controller

	// per user turn
	// - render map
	viewMap.Render(m)
	// - get resources
	// - per robot
	//   - move
	//   - ability
	//   - combat


	// base(s) actions

	// end turn?
	//TODO:fmt.Scanln()
}
