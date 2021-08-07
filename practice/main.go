package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	for i, _ := range colors {
		fmt.Println(i)
	}

	i := 1
	switch i {
	case 1:
		fmt.Println("hello")
		fallthrough
	case 2:
		fmt.Println("Hi")
	}
}
