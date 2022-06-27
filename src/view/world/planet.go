package world

import (
	"fmt"

	"github.com/honggim/robots/src/model/world"
)

func Render(p *world.Planet) {
	fmt.Println(p.Name)
	fmt.Println("- length: ", p.Length())
	fmt.Println("-  width: ", p.Width())

	fmt.Println(renderTiles(p.Tiles))
}

func renderTiles(tiles [][]*world.Tile) string {
	output := ""

	for i := 0; i < len(tiles); i++ {
		output += renderRow(tiles[i])
		output += "\n"
	}

	return output
}

func renderRow(row []*world.Tile) string {
	output := ""
	for i := 0; i < len(row); i++ {
		output += renderTile(row[i])
	}

	return output
}

func renderTile(t *world.Tile) string {
	return fmt.Sprintf("[%3d/%3d]", t.Current, t.Total)
}
