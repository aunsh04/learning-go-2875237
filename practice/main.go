package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("sum:", intSum)

	f1, f2 := 10.2, 15.6
	floatSum := f1 + f2
	fmt.Println("Float sum:", math.Round(floatSum))

	radius := 15.2
	circumference := 2 * math.Pi * radius
	fmt.Printf("Circumference: %.2f\n", circumference)

}
