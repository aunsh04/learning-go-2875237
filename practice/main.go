package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var colors = [3]string{"h", "k", "l"}
	fmt.Println(colors)
	fmt.Println(len(colors))

	var colors1 [2]string
	colors1[0] = "j"
	colors1[1] = "y"

	fmt.Println(colors1)

	var p = 42
	k := &p
	fmt.Println(*k)
}
