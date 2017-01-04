package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, My name is ", p.Name)
}

type Android struct {
	Person //anonymous fields
	Modle  string
}

func main() {
	a := new(Android)
	a.Talk()
	a.Person.Talk()
}
