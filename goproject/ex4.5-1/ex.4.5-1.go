package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = a
	// d = a * b 서로 다른 타입간 연산 불가

	var e int64 = 7
	// f := a * e 같은 int 이지만 서로 크기가 다르므로 연산 불가

	fmt.Println(a, b, c, e)
}