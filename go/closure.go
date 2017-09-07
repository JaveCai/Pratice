package main

import (
	"fmt"
)

func main() {

	x := 1

	y := func() {
		fmt.Println("x:", x)
		x++
	}

	for i := 0; i < 10; i++ {
		y()
	}
}
