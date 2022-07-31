package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

type User struct {
	Name string
	ID   string
	Age  int
}

type EmbeddedUser struct {
	Name string
}

type VIPUser struct {
	UserInfo     User
	VIPLevel     int
	Price        int
	EmbeddedUser // Embedded 처럼 사용하려면 Field 명을 생략하면 됨.
}

type TestUser struct {
	A int8
	B int
	C int8
	D int
	E int8
}

type TestUser2 struct {
	A int8
	B int8
	C int8
	D int
	E int
}

func main() {
	var house House
	house.Address = "서울시 송파구"
	house.Size = 10
	house.Price = 9.8
	house.Type = "아파트"

	var house2 House = House{"경기도 고양시", 30, 5.0, "단독주택"}

	var house3 House = House{Address: "경기도 고양시", Size: 30, Price: 5.0, Type: "단독주택"}

	fmt.Println("주소 : ", house.Address)
	fmt.Println("크기 : ", house.Size)
	fmt.Println("가격 : ", house.Price)
	fmt.Println("타입 : ", house.Type)

	fmt.Println("주소 : ", house2.Address)
	fmt.Println("크기 : ", house2.Size)
	fmt.Println("가격 : ", house2.Price)
	fmt.Println("타입 : ", house2.Type)

	fmt.Println("주소 : ", house3.Address)
	fmt.Println("크기 : ", house3.Size)
	fmt.Println("가격 : ", house3.Price)
	fmt.Println("타입 : ", house3.Type)

	fmt.Println("========================================")

	user := User{Name: "송하나", ID: "hana", Age: 13}

	vip := VIPUser{
		UserInfo:     user,
		VIPLevel:     3,
		Price:        250,
		EmbeddedUser: EmbeddedUser{Name: "TEST USER"},
	}

	fmt.Printf("VIP 유저: %s | ID: %s | 나이 : %d | VIP Level : %d | VIP 가격: %d 만원 (Embedded User Name : %s)\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price,
		vip.Name, // Embedded 처럼 사용하려면 Field 명을 생략하면 됨.
	)

	fmt.Println("========================================")

	fmt.Printf("Sturucture Size : %d \n", unsafe.Sizeof(vip))

	fmt.Println("========================================")

	testUser := TestUser{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(testUser)) // print 40

	fmt.Println("Memory Padding")

	testUser2 := TestUser2{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(testUser2)) // print 24
}
