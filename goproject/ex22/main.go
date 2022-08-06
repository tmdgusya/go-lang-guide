package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()

	e4 := v.PushBack(4)
	e1 := v.PushFront(1)

	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, " ")
	}

	fmt.Println()

	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value, " ")
	}

	m := make(map[int]string)

	m[10] = "test"

	fmt.Println(m[10])

	for k, v := range m {
		fmt.Println(k, v)
	}
}
