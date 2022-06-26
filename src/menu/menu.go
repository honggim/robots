package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/honggim/robots/src/model"
	"github.com/honggim/robots/src/model/robot"
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

/*
func (m *Menu) GetRobots() *model.Robot {
	for i := 0; i < 3; i++ { //TODO: # of bots as game option
	}
}
*/

func (m *Menu) GetMockRobots() []*robot.Robot {
	bots := make([]*robot.Robot, 3)

	bots[0] = robot.NewRobot("optimus prime")
	bots[1] = robot.NewRobot("bumble bee")
	bots[2] = robot.NewRobot("ironhide")

	return bots
}
