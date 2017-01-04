package main

import (
	"fmt"
)

func average(xs []float64) float64 {

	total := 0.0
	for _, v := range xs {
		total += v
	}
	//	panic("Not Implemented")
	return (total / float64(len(xs)))
}

func main() {
	xs := []float64{98, 93, 43, 23, 96}
	fmt.Println(average(xs))

}
