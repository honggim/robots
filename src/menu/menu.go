package menu

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/honggim/robots/src/model/robot"
	"github.com/honggim/robots/src/model/world"
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

func (m *Menu) File() error {

	return nil
}

//TODO:func (m *Menu) Save(filename string, data any) error {
func (m *Menu) Save(filename string, data interface{}) error {
	//TODO: raw, err := json.MarshalIndent(data, "", "  ")
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//TODO: saves entire to memory to file; need to address if too big for memory
	return ioutil.WriteFile(filename, raw, 0644)
}

func (m *Menu) Load(filename string, data interface{}) {
	raw, err := ioutil.ReadFile(filename)
	check(err)

	err = json.Unmarshal(raw, &data)
	check(err)
}

func (m *Menu) GetMapOptions() *world.PlanetOptions {
	fmt.Print("name of game? ")
	name := m.readString()

	fmt.Print("length of map? ")
	length := m.readInt()

	fmt.Print("width of map? ")
	width := m.readInt()

	fmt.Print("max energon? ")
	energon := 100

	return &world.PlanetOptions{
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

func (m *Menu) GetMockMapOptions() *world.PlanetOptions {
	return &world.PlanetOptions{
		Name:    "mock name",
		Length:  3,
		Width:   4,
		Energon: 100,
	}
}

func (m *Menu) GetMockRobots() []*robot.Robot {
	bots := make([]*robot.Robot, 3)

	bots[0] = robot.NewRobot("optimus prime")
	bots[1] = robot.NewRobot("bumble bee")
	bots[2] = robot.NewRobot("ironhide")

	return bots
}
