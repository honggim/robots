package main

import (
	"fmt"

	"github.com/honggim/robots/src/inputs"
	"github.com/honggim/robots/src/maps"
)

func main() {
	// create & login user

	// get game options
	// - map: size, features, i.e civ
	m := maps.Generate(inputs.GetInputs())
	fmt.Println(m)
	// - pick core robots


	// per user turn
	// - see tiles
	// - get resources
	// - per robot
	//   - move
	//   - ability
	//   - combat


	// base(s) actions

	// end turn
}
