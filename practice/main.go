package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Weigh()

}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}

// Speak is how the dog speaks
func (d Dog) Weigh() {
	fmt.Println(d.Weight)
}
