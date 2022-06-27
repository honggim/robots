package world

type Planet struct {
	Name   string
	Tiles  [][]*Tile
}

type PlanetOptions struct {
	Name    string
	Length  int
	Width   int
	Energon int
}

func NewPlanet(opts *PlanetOptions) *Planet {
	//TODO: if length or width < 1

	factory := NewFactory(opts.Energon)

	tiles := make([][]*Tile, opts.Length)
	for i := 0; i < opts.Length; i++ {
		tiles[i] = factory.NewTiles(opts.Width)
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
