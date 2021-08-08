package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")

	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, "Hi My name is Aunsh")
	checkError(err)
	fmt.Println(length)
	defer file.Close()
	defer readFile("./fromString.txt")

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from:", string(data))
}
