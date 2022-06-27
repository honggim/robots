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
	//TODO: use enum/iota
}

type TileFactory struct {
	maxEnergon int
}

func NewFactory(maxEnergon int) *TileFactory {
	// needed to randomize rand.Intn()
	rand.Seed(time.Now().UnixNano())

	return &TileFactory{
		maxEnergon: maxEnergon,
	}
}

func (tf *TileFactory) NewTiles(num int) []*Tile {
	tiles := make([]*Tile, num)

	for i := 0; i < num; i++ {
		energon := rand.Intn(tf.maxEnergon)
		tiles[i] = &Tile{
			Current: energon,
			Total:   energon,
		}
	}

	return tiles
}
