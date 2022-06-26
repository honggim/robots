package main

import (
	"fmt"

	"github.com/honggim/robots/src/maps"
)

func main() {
	fmt.Println("main(): start")
	defer fmt.Println("main(): end")

	//maps.Generate(4,5)
	hash := maps.Generate("first map", 3, 6)
	fmt.Println(hash)
}
