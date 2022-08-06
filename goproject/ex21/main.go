package main

import (
	"fmt"
	"os"
)

type opFunc func(a, b int) int

// Functional Programming
func getOpreator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	}

	if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	}

	return nil
}

func main() {
	result := sum(1, 2, 3, 4, 5)

	fmt.Printf("Sum Result %d\n", result)

	printf(1, true)

	f, err := os.Create("test.txt")

	if err != nil {
		fmt.Println("Failed to Create File")
		return
	}

	defer fmt.Println("반드시 호출 됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Hello World 를 씁니다.")
	fmt.Fprintln(f, "Hello World")

	fmt.Println("*******************************")

	operator := getOpreator("*")

	fmt.Println(operator(3, 4))

	fmt.Println("*******************************")

	i := 0

	f1 := func() {
		i += 20
	}

	f1()

	fmt.Println(i)
}

// mutiple variables
func sum(nums ...int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}
	return sum
}

func printf(args ...interface{}) {
	for _, arg := range args {
		switch i := arg.(type) {
		case bool:
			fmt.Println(i)
		case float64:
			fmt.Println(i)
		case int:
			fmt.Println(i)
		default:
			fmt.Print("1231312")
		}
	}
}
