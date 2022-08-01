package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	var slice2 = []int{1, 5: 2, 10: 3}

	fmt.Println(slice)
	fmt.Println(slice2)

	fmt.Println("=============================")

	var slice3 = make([]int, 4, 5)
	fmt.Println(slice3)

	fmt.Println("=============================")

	slice4 := append(slice3, 5) // new instance
	fmt.Println(slice4)
	fmt.Println(slice3)

	slice3[0] = 100

	fmt.Println(slice4)
	fmt.Println(slice3)

	fmt.Println("=============================")

	array := [5]int{1, 2, 3, 4, 5}
	slice6 := array[1:2]

	fmt.Println(slice6)

	array[1] = 100

	fmt.Println(slice6)

	fmt.Println("=============================")

	slice_1 := []int{1, 2, 3, 4, 5}
	slice_2 := make([]int, 3, 10)
	slice_3 := make([]int, 10)

	cnt1 := copy(slice_2, slice_1)
	cnt2 := copy(slice_3, slice_1)

	fmt.Println(cnt1, slice_2)
	fmt.Println(cnt2, slice_3)

	fmt.Println("=============================")

	slice_del := append(slice_1[:1], slice_1[2:]...)
	fmt.Println(slice_del)
}
