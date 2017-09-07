package main

import (
	"fmt"
)

func main() {

	x := 1

	y := func() {
		fmt.Println("x:", x)
		x++ //the same action as a static valuable in C/C++
	}

	for i := 0; i < 10; i++ {
		y() //1,2,3...10
	}
}
