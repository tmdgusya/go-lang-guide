package main

import "fmt"

func main() {
	var t [5]float32 = [5]float32{24.0, 25.9, 27.8, 26.9, 26.2}

	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

	// array size assign 5
	nums := [...]int{10, 20, 30, 40, 50}

	nums[2] = 300

	for i := 0; i < 5; i++ {
		fmt.Println(nums[i])
	}

	// range function
	for index, value := range nums {
		fmt.Printf("Index : %d , Value: %d \n", index, value)
	}

	// array 가 메모리 지역성 때문에 검색속도가 월등히 빠름.

	// array value copy

	fmt.Println("=========ARRAY COPY TEST=========")

	target1 := [...]int{10, 20, 30, 40, 50}
	target2 := [...]int{60, 70, 80, 90, 100}

	target2 = target1

	target1[0] = 100

	for i, v := range target2 {
		fmt.Printf("Index : %d , Value: %d \n", i, v) // 0번째값이 10으로 나오는 것을 보아 값을 copy 하는걸로 보인다.
	}

	fmt.Println("=========MULTIPLE ARRAY=========")

	b := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	fmt.Println("Length : ", len(b))

	for _, innerArray := range b {
		for _, innerValue := range innerArray {
			fmt.Print(innerValue, " ")
		}
	}
}
