package model

import (
	"math/rand"
	"time"
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
	Tiles  [][]Tile
}

type MapOptions struct {
	Name    string
	Length  int
	Width   int
	Energon int
}

func NewMap(opts *MapOptions) *Map {
	//TODO: if length or width < 1

	// needed to randomize rand.Intn()
	rand.Seed(time.Now().UnixNano())

	tiles := make([][]Tile, opts.Length)

	for l := 0; l < opts.Length; l++ {
		tiles[l] = make([]Tile, opts.Width)
		for w := 0; w < opts.Width; w++ {
			tiles[l][w] = randomTile(opts.Energon)
		}
	}

	return &Map{
		Name:   opts.Name,
		Tiles:  tiles,
	}
}

func (m *Map) Length() int {
	return len(m.Tiles)
}

func (m *Map) Width() int {
	return len(m.Tiles[0])
}
