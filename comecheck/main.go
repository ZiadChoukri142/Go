package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "galaxy" || arguments[i] == "01" || arguments[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
		}
	}
}
