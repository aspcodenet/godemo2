package loopdemos

import "fmt"

/// access modifiers PUBLIC, PRIVATE, PROTECTED

var NrOfRuns int

func Demo2() { // public
	NrOfRuns++
	fmt.Println("hej")
	demo1()
}

func demo1() { // private
	fmt.Println("hej")
}
