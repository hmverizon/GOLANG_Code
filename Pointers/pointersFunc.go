package main

import (
	"fmt"
)

func zeo(ptr *int) {
	*ptr += 2
}

func main() {
	x := 2
	zeo(&x)
	fmt.Println(x)
	anotherWayofPointer()
}

func anotherWayofPointer() {
	x := new(int)

	zeo(x)
	fmt.Println(*x)
}
