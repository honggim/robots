package main

import (
	"fmt"

	"encoding/json"
	"io/ioutil"

	"github.com/honggim/robots/src/menu"
	"github.com/honggim/robots/src/model/robot"
	"github.com/honggim/robots/src/model/world"
	viewWorld "github.com/honggim/robots/src/view/world"
)

var m *menu.Menu
var filename string
var planet *world.Planet
var bots []*robot.Robot

func init() {
	m = menu.NewMenu()

	// login screen
	//menu.Login()

	// new / save / load game
	filename = "savefile.txt"
	//m.File()
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("type: %T, raw: %s", raw, raw)
	fmt.Printf("type: %T\n", raw)
	fmt.Println(string(raw))
	var p world.Planet
	if err = json.Unmarshal(raw, &p); err != nil {
		fmt.Println("err:", err)
		fmt.Println("planet:", &p)
		panic(err)
	}
	viewWorld.Render(&p)

	// - map: size, features, i.e civ
	//planet = world.NewPlanet(m.GetMockMapOptions())

	// - create core robots
	/*
	bots := m.GetMockRobots()
	for i := 0; i < len(bots); i++ {
		bots[i].Render()
	}
	*/
}

func main() {
	//TODO: move into controller

	// per user turn
	// - render map
	//viewWorld.Render(planet)
	// - get resources
	// - per robot
	//   - move
	//   - ability
	//   - combat


	// base(s) actions

	// end turn?
	//m.Save("savefile.txt", planet)
	//m.Save(filename, planet)
	fmt.Scanln()
}
