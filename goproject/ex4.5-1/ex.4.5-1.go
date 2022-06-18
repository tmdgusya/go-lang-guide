package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	var c int = a
	d := float64(a * c)

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}