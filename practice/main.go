package main

import (
	"fmt"
)

func main() {
	anInt := 10
	p := &anInt
	fmt.Println("Value of p:", *p)
	fmt.Println("Pointers")

	value1 := 50
	p1 := &value1
	fmt.Println(p1)
	fmt.Println(*p1)

	*p1 = *p1 * 20
	fmt.Println(*p1)
	fmt.Println(value1)
}
