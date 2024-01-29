package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	result := os.Args[0]
	lastSlashIndex := -1
	for i := 0; i < len(result); i++ {
		if result[i] == '/' {
			lastSlashIndex = i
		}
	}

	if lastSlashIndex != -1 {
		programName := result[lastSlashIndex+1:]
		for _, r := range programName {
			z01.PrintRune(r)
		}
	} else {
		for _, r := range result {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
