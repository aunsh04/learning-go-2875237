package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	poodle := Dog{"Dalmatian", 10}
	fmt.Println(poodle.Weight)
}

type Dog struct {
	Breed  string
	Weight int
}
