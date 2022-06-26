package maps

/*
import (
	"fmt"
)
*/

type Map struct {
	name   string
	length int
	width  int
	tiles  [][]Tile
}

type Tile struct {
}

func Generate(name string, length, width int) Map {
	m := Map{
		name: name,
		length: length,
		width: width,
	}

	return m
}
