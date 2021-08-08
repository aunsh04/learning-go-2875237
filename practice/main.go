package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	sum := addValues(10, 10)
	fmt.Println(sum)
	newSum, newSum1 := addMultipleValues(10, 10, 10)
	fmt.Println(newSum)
	fmt.Println(newSum1)
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addMultipleValues(values ...int) (int, int) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, 9
}
