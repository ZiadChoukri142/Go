package main

import "github.com/01-edu/z01"

type Door struct {
	Door  string
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr(" Door Closing...")
	ptrDoor.state = false
	return true
}

func IsDoorOpen(Door Door) bool {
	PrintStr(" is the Door opened ?")
	return Door.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr(" is the Door closed ?")
	return !ptrDoor.state
}

func main() {
	door := &Door{}

	CloseDoor(door)
	if IsDoorClose(door) {
		PrintStr(" The door is closed.\n")
	} else {
		PrintStr(" The door is open.\n")
	}
}
