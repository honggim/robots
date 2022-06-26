package main

import (
	"fmt"

	"github.com/honggim/robots/src/menu"
	"github.com/honggim/robots/src/models"
	"github.com/honggim/robots/src/views"
)

var m *models.Map

func init() {
	menu := menu.NewMenu()

	// login screen
	//menu.Login()

	//TODO: menu for new game, save, load

	// - map: size, features, i.e civ
	//m = models.GenerateMap(menu.GetMapOptions())
	m = models.GenerateMap(menu.GetMockMapOptions())

	// - pick core robots
}

func main() {
	//TODO: move into controller

	// per user turn
	// - render map
	views.Render(m)
	// - get resources
	// - per robot
	//   - move
	//   - ability
	//   - combat


	// base(s) actions

	// end turn?
	//TODO:fmt.Scanln()
}
