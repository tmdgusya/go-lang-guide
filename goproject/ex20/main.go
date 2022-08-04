package main

import "fmt"

type DuckInterface interface {
	Fly()
	Walk(distance int) int
}

type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Quack")
}

func (d Duck) Walk(distance int) int {
	return distance
}

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float32:
		fmt.Printf("v is int %f\n", float32(t))
	default:
		fmt.Println("Not Supported!")
	}
}

func main() {
	duck := &Duck{}

	var interDuck DuckInterface

	interDuck = duck

	interDuck.Fly()
	fmt.Println(interDuck.Walk(10))
}
