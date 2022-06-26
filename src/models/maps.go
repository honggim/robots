package models

import (
	"math/rand"
)

type Tile struct {
	Current int
	Total   int
	//TODO: speed   int
	//TODO: type    string
}

//TODO: use enum/iota

func randomTile(maxEnergon int) Tile {
	energon := rand.Intn(maxEnergon)
	tile := Tile{
		Current: energon,
		Total:   energon,
	}

	return tile
}

type Map struct {
	Name   string
	Length int
	Width  int
	Tiles  [][]Tile
}

func GenerateMap(name string, length, width, maxEnergon int) *Map {
	tiles := make([][]Tile, length)

	for l := 0; l < length; l++ {
		tiles[l] = make([]Tile, width)
		for w := 0; w < width; w++ {
			tiles[l][w] = randomTile(maxEnergon)
		}
	}

	return &Map{
		Name:   name,
		Length: length,
		Width:  width,
		Tiles:  tiles,
	}
}
