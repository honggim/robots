package main

import (
	"fmt"

	//"github.com/honggim/robots/src/config/inputs"
	"github.com/honggim/robots/src/menu"
	"github.com/honggim/robots/src/models"
)

var m models.Map

func init() {
	menu := menu.NewMenu()

	// login screen
	menu.Login()

	//TODO: menu for new game, save, load

	// - map: size, features, i.e civ
	m := models.GenerateMap(menu.GetMapOptions())
	fmt.Println(m)

	// - pick core robots
}

func main() {
	//TODO: move into controller
	// per user turn
	// - see tiles
	// - get resources
	// - per robot
	//   - move
	//   - ability
	//   - combat


	// base(s) actions

	// end turn?
	//TODO:fmt.Scanln()
}
