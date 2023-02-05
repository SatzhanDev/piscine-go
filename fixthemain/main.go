// package piscine

// func PrintStr(s string) {
// 	for _, r := range s {
// 		z01.PrintRune(r)
// 	}
// }

// func CloseDoor(ptrDoor *Door) bool {
// 	PrintStr("Door Closing...")
// 	state = CLOSE
// 	return true
// }

// func OpenDoor(ptrDoor *Door) bool {
// 	PrintStr("Door Closing...")
// 	state = CLOSE
// 	return true
// }

// func IsDoorOpen(ptrDoor *Door) {
// 	PrintStr("is the Door opened ?")
// }

// func IsDoorClose(ptrDoor *Door)  {
// 	PrintStr("is the Door closed ?")
// }

// func main() {
// 	door := &Door{}

// 	OpenDoor(door)
// 	if IsDoorClose(door) {
// 		OpenDoor(door)
// 	}
// 	if IsDoorOpen(door) {
// 		CloseDoor(door)
// 	}
// 	if door.state == OPEN {
// 		CloseDoor(door)
// 	}
// }

package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

func PrintStr(str string) {
	for _, s := range str {
		z01.PrintRune(s)
	}
	z01.PrintRune(10)
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = true
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = false
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state == true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == false
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == true {
		CloseDoor(door)
	}
}
