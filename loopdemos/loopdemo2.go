package loopdemos

import "fmt"

func LoopFor1() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)
}

func LoopWhile() {
	n := 1
	for n < 5 { // while n < 5
		n *= 2
	}
	fmt.Println(n)
}

func LoopInfinite() {
	sum := 0
	for { // while true
		sum++
		if sum > 100 {
			break
		}
	}
}

func LoopRange() {
	strings := []string{"hello", "world"}
	for index, str := range strings {
		fmt.Println(index, str) // 0, "hello" 1 "world"
	}

}

func LoopRange2() {
	strings := []string{"hello", "world"}
	for _, str := range strings {
		fmt.Println(str) // "hello" "world"
	}

}
