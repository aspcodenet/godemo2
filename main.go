package main

import (
	"fmt"
	"main/loopdemos"
)

func test() {

}

func add(x int, y int) (result int) {
	result = x + y
	return
}

func main() {
	d := 12
	res := add(d, 19)
	fmt.Println(res)

	// var a = new Player()
	// a.Calcuilate()

	loopdemos.Demo2()
	loopdemos.NrOfRuns++
	hejhej()

	i := theCoolNumber

	var i3 int
	i3 = 12

	e := 14 // WALRUS = deklarera och tilldela i en och samma

	// a = $"Hej <name>${name}</name> dfssdfsdf"
	name := "Stefan"
	fmt.Printf("Hej %v  <name>%v</name> %v %v\n", i, name, i3, e)
}
