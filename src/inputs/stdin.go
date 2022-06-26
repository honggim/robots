package inputs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
//TODO: implement interface
type Inputs interface {
	GetInputs() (string, int, int)
}
*/

func GetInputs() (name string, length, width int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("name of game? ")
	rawName, err := reader.ReadString('\n')
	check(err)
	name = string(rawName)

	fmt.Print("length of map? ")
	rawLength, err := reader.ReadString('\n')
	length, err = strconv.Atoi(string(rawLength[:len(rawLength)-1]))
	check(err)

	fmt.Print("width of map? ")
	rawWidth, err := reader.ReadString('\n')
	width, err = strconv.Atoi(string(rawWidth[:len(rawWidth)-1]))
	check(err)

	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetMockInputs() (string, int, int) {
	return "mock name", 3, 6
}
