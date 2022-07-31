package main

import "fmt"

func getMyAge() (int, bool) {
	return 33, true
}

func main() {
	// 특정 if 문에서만 사용해야 한다면 이렇게 사용하는게 좋은 듯 함.
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are beautiful", age)
	} else {
		fmt.Println("Error") // bb
	}

	// fmt.Println("Your age is", age) 사용불가
}
