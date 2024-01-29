package main

import "github.com/01-edu/z01"

const value = "x = 42, y = 21\n"

type point struct {
	x int
	y int
}

func SetPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	SetPoint(points)
	for _, r := range value {
		z01.PrintRune(r)
	}
}
