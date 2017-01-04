package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}
type circle2 struct {
	x, y, r float64
}

func (circle2 c) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	x := new(circle2)
	x.r = 6
	x.y = 8
	x.x = 2
	fmt.Println(x.r, x.y, x.x)
	fmt.Println(c.area())

}
