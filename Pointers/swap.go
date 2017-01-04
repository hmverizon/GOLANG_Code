package main

import "fmt"

func swap(ptr1 *int, ptr2 *int) {
	temp := new(int)
	*temp = *ptr1
	*ptr1 = *ptr2
	*ptr2 = *temp
}

func main() {
	//	a := new(int)
	//b := new(int)
	a := 9
	b := 8
	swap(&a, &b)

	fmt.Println(a, b)
}
