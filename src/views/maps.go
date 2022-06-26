package maps

import (
	"fmt"

	"github.com/honggim/robots/src/models"
)

func Render(m *models.Map) {
	line := fmt.Sprintf("name: %s [%d x %d] (l x w)", m.Name, m.Length, m.Width)
	fmt.Println(line)
}
