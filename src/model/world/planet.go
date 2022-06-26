package world

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

//TODO: rename to "Planet"
type Planet struct {
	Name   string
	Tiles  [][]Tile //TODO: chnage to pointers?
}

type PlanetOptions struct {
	Name    string
	Length  int
	Width   int
	Energon int
}

func NewPlanet(opts *PlanetOptions) *Planet {
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

	return &Planet{
		Name:   opts.Name,
		Tiles:  tiles,
	}
}

func (p *Planet) Length() int {
	return len(p.Tiles)
}

func (p *Planet) Width() int {
	return len(p.Tiles[0])
}
