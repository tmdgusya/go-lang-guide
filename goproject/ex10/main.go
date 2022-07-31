package main

import "fmt"

func returnAnyIntegerValue() int {
	return 3
}

type ColorType int

const (
	Red ColorType = iota
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"
	}
}

func main() {

	a := 3

	switch a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3?!")
		fallthrough
	default:
		fmt.Println("a is bigger than 3")
	}

	// switch 도 이런식으로 가능하다.
	switch value := returnAnyIntegerValue(); value {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is bigger than 3")
	}

	fmt.Println(colorToString(Red))

}
