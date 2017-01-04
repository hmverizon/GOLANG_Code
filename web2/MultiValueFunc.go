package main

import (
	"fmt"
)

func mVal() (int, int) {
	return 5, 6
}

func main() {

	x, y := mVal()
	fmt.Println(x, y)
}
