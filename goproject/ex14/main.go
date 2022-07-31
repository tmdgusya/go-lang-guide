package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

func ChangeData(arg Data) {
	// memory address is 1608 byte size
	arg.value = 999
	arg.data[100] = 999
}

func ChangeRealData(arg *Data) {
	// memory address is only 8byte size because it's just integer number
	arg.value = 999
	arg.data[100] = 999
}

type HeapUser struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *HeapUser {
	var u = HeapUser{Name: name, Age: age}
	return &u // Do not free memory because of escape-analysis
}

func main() {
	// pointer can save memory address

	// variable 'p' is save memory address of integer type
	var p *int
	var a int = 0

	if p == nil {
		fmt.Println("Any Memory Address doesn't assigned to 'p'")
	}

	// A give memory address of a to B
	p = &a

	if p != nil {
		fmt.Println("An Memory Address do assigned to 'p'")
	}

	fmt.Printf("A(p) memory Address : %d \n", p)
	fmt.Printf("A(*p) memory Address : %d \n", *p)
	fmt.Printf("A memory Address : %d \n", &a)
	fmt.Printf("A* memory Address : %d \n", a)

	fmt.Println("a == p ", (a == *p))

	// p is point to memory address of 'a'

	*p = 20

	fmt.Printf("A(p) memory Address : %d \n", p)
	fmt.Printf("A(*p) memory Address : %d \n", *p)
	fmt.Printf("A memory Address : %d \n", &a)
	fmt.Printf("A* memory Address : %d \n", a)

	fmt.Println("========================================")

	var data Data

	ChangeData(data) // not changed value because it is just copied that
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	fmt.Println("========================================")

	ChangeRealData(&data) // changed value because do give memory address of mine to function through arguments
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

	fmt.Println("========================================")

	var instance1 *Data = &Data{value: 1, data: [200]int{10}}

	var instance2 *Data = instance1
	var instance3 *Data = instance1

	fmt.Printf("Instance1 Address Value: %d \n", &instance1)
	fmt.Printf("Instance2 Address Value: %d \n", &instance2)
	fmt.Printf("Instance3 Address Value: %d \n", &instance3)

	instance2.value = 3

	fmt.Println("instance1.value == instance2.value : ", (instance1.value == instance2.value)) // true
	fmt.Println("instance3.value == instance2.value : ", (instance3.value == instance2.value))
	fmt.Printf("Instance1 Value: %d \n", instance1.value)

	// If you use 'new' when create new instance, then you can't assign specify value to variable
	var instance4 = new(Data)

	fmt.Printf("Instance4 Value: %d \n", instance4.value)
	fmt.Println("instance1.value == instance4.value : ", (instance1.value == instance4.value)) // false

	fmt.Println("========================================")

	userPointer := NewUser("AAA", 23)

	fmt.Println(userPointer)
}
