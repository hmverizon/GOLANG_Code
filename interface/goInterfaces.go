package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}

func totalarea(Shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
}
func main() {

	fmt.Println(totalarea(&c, &r))
}
