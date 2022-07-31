package main

import "fmt"

// 대문자로 적으면 Public 임. 현재 상수는 외부 Package 로 공개되는 상수이다.
const ConstValue int = 10

// const C int = 10

// var b int = C * 20
// C = 20 -> Error 상수는 다시 대입 불가
// fmt.Println(&C) // 상수에는 메모리 주솟값 접근이 불가능함.

// create enum using iota
const (
	RED   int = iota // 0
	BLUE  int = iota // 1
	GREEN int = iota // 2
)

// 컴퓨터에서 상수는 리터럴로 변환되어 실행파일에 쓰인다. 띠리사 상수 표현식에 CPU 자원을 사용하지 않는다. 동적 할당메모리 영역 또한 사용하지 않음.

func main() {
	const PI1 float64 = 3.141592653589798238
	var PI2 float64 = 3.141592653589798238

	PI2 = 4

	fmt.Printf("원주율 : %f\n", PI1)
	fmt.Printf("원주율 : %f\n", PI2)
}
