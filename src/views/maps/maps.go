package maps

import (
	"fmt"

	"github.com/honggim/robots/src/models"
)

func Render(m *models.Map) {
	fmt.Println(m.Name)
	fmt.Println("- length: ", m.Length())
	fmt.Println("-  width: ", m.Width())

	for i := 0; i < m.Length(); i++ {
		fmt.Println(renderRow(m.Tiles[i]))
	}
}

func renderRow(row []models.Tile) string {
	output := ""
	for i := 0; i < len(row); i++ {
		output += renderTile(row[i])
	}

	return output
}

func renderTile(t models.Tile) string {
	return fmt.Sprintf("[%3d/%3d]", t.Current, t.Total)
}
