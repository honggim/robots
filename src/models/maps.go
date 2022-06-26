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

func randomTile() Tile {
	energon := rand.Intn(10)
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

func GenerateMap(name string, length, width int) *Map {
	tiles := make([][]Tile, length)

	for l := 0; l < length; l++ {
		tiles[l] = make([]Tile, width)
		for w := 0; w < width; w++ {
			tiles[l][w] = randomTile()
		}
	}

	return &Map{
		Name:   name,
		Length: length,
		Width:  width,
		Tiles:  tiles,
	}
}
