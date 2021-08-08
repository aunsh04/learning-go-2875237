package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	in1 := acceptInput("Value 1")
	in2 := acceptInput("Value 2")

	switch operation := acceptOperation(); operation {

	case "+":
		fmt.Println("Adding the numbers")
	case "-":
		fmt.Println("Subtracting the numbers")
	default:
		panic("Invalid operation")
	}

	fmt.Println(in1 + in2)
	// fmt.Print("Value 1: ")
	// input1, _ := reader.ReadString('\n')
	// float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Value 1 must be a number")
	// }

	// fmt.Print("Value 2: ")
	// input2, _ := reader.ReadString('\n')
	// float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("Value 2 must be a number")
	// }

	// sum := float1 + float2
	// sum = math.Round(sum*100) / 100
	// fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)
}

func acceptInput(prompt string) float64 {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value must be a number")
	}
	return float1
}

func acceptOperation() string {
	fmt.Print("Enter operation")
	reader := bufio.NewReader(os.Stdin)
	operation, _ := reader.ReadString('\n')
	return strings.TrimSpace(operation)
}
