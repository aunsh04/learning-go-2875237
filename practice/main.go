package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[2 : len(colors)-1])
	fmt.Println(colors)

	var colors1 = []string{"Red", "Green", "Blue"}

	sort.Strings(colors1)
	fmt.Println(colors1)

}
