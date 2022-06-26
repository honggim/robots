package maps

import (
	//"fmt"
	"math/rand"
)

type Tile struct {
	current int
	total   int
	speed   int
	//TODO:type    string
}

//TODO: use enum/iota

func randomTile() Tile {
	energon := rand.Intn(10)
	tile := Tile{
		current: energon,
		total:   energon,
		speed:   rand.Intn(3),
	}

	return tile
}

type Map struct {
	name   string
	length int
	width  int
	tiles  [][]Tile
}

/*
func (m *Map) String() string {
	//TODO: print tiles
	return fmt.Sprintf("name: %s\nlength: %d\nwidth: %d", m.name, m.length, m.width)
}
*/

func Generate(name string, length, width int) *Map {
	tiles := make([][]Tile, length)

	for l := 0; l < length; l++ {
		tiles[l] = make([]Tile, width)
		for w := 0; w < width; w++ {
			tiles[l][w] = randomTile()
		}
	}

	return &Map{
		name:   name,
		length: length,
		width:  width,
		tiles:  tiles,
	}
}
