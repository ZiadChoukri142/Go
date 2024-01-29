package main

import (
	"os"

	"github.com/01-edu/z01"
)

func findMinArg(args []string) int {
	minIndex := 0
	for i := 1; i < len(args); i++ {
		if args[i] < args[minIndex] {
			minIndex = i
		}
	}
	return minIndex
}

func main() {
	args := os.Args[1:]

	for len(args) > 0 {
		minIndex := findMinArg(args)
		for _, r := range args[minIndex] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')

		args = append(args[:minIndex], args[minIndex+1:]...)
	}
}
