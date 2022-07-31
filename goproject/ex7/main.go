package main

import "fmt"

// argument 는 주입받은 값을 copy 해서 사용함. 원본 값을 사용하지 않음.
func Add(a int, b int) int {
	return a + b
}

// Function can return multi value
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func Divide2(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // 명시적으로 적어주지 않아도 리턴됨.
	}
	result = a / b
	success = true
	return
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)

	c, success := Divide(3, 6)
	fmt.Println(c, success)

	// d, success := Divide2(3, 6) 요게 가능하긴 함. 꼭 원하는 걸로 받으라는건 아닌듯함.
	result, success := Divide2(3, 6)
	fmt.Println(result, success)
}
