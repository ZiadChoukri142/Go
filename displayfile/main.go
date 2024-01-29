package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	length := len(os.Args)
	if length < 2 {
		fmt.Println("File name missing")
		return
	}

	if length > 2 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := os.Args[1]

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print("Error reading file:", err)
		return
	}

	fmt.Print(string(content))
}
