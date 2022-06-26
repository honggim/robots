package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/honggim/robots/src/model"
)

type Menu struct {
	reader *bufio.Reader
}

func NewMenu() *Menu {
	return &Menu{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (m *Menu) Login() {
	fmt.Println("Login Screen")
	fmt.Println("============")

	fmt.Print("username: ")
	name := m.readString()
	msg := fmt.Sprintf("Welcome %s!", name)
	fmt.Println(msg)
}

func (m *Menu) GetMapOptions() *model.MapOptions {
	fmt.Print("name of game? ")
	name := m.readString()

	fmt.Print("length of map? ")
	length := m.readInt()

	fmt.Print("width of map? ")
	width := m.readInt()

	fmt.Print("max energon? ")
	energon := 100

	return &model.MapOptions{
		Name:    name,
		Length:  length,
		Width:   width,
		Energon: energon,
	}
}

func (m *Menu) readInt() int {
	raw := m.readString()
	num, err := strconv.Atoi(raw)
	check(err)

	return num
}

func (m *Menu) readString() string {
	raw, err := m.reader.ReadString('\n')
	check(err)

	return raw[:len(raw)-1]
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (m *Menu) GetMockMapOptions() *model.MapOptions {
	return &model.MapOptions{
		Name:    "mock name",
		Length:  3,
		Width:   4,
		Energon: 100,
	}
}
