package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {

	poet := `
		죽는 날 까지 하늘을 우러러 
		한점 부끄럼이 없기를,
		잎새에 이는 바람에도
		나는 괴로워했다.
	`

	fmt.Println(poet)

	fmt.Println("========================================")

	// defaut uniset is utf-8 in golang

	var char rune = '한'

	fmt.Printf("%T\n", char) // print type of char
	fmt.Println(char)        // print value of char
	fmt.Printf("%c\n", char) // print value of char to string

	fmt.Println("========================================")

	str1 := "가나다라마"
	str2 := "abcde"

	fmt.Printf("len(str1) = %d\n", len(str1)) // 한글은 3 바이트 이므로.
	fmt.Printf("len(str2) = %d\n", len(str2))

	str3 := "Hello 월드"
	runesFromStr3 := []rune(str3) // rune is set of chars

	fmt.Printf("len(str3) = %d\n", len(str3))
	fmt.Printf("len(runesFromStr3) = %d\n", len(runesFromStr3))

	fmt.Println("========================================")

	str4 := "Hello 월드"

	for i := 0; i < len(str4); i++ {
		fmt.Printf("타입 : %T | 값: %d | 문자값: %c\n", str4[i], str4[i], str4[i])
	}

	fmt.Println("========================================")

	runesFromStr4 := []rune(str4) // unncessary use memory

	for i := 0; i < len(runesFromStr4); i++ {
		fmt.Printf("타입 : %T | 값: %d | 문자값: %c\n", runesFromStr4[i], runesFromStr4[i], runesFromStr4[i])
	}

	fmt.Println("========================================")

	for _, v := range str4 {
		fmt.Printf("타입 : %T | 값: %d | 문자값: %c\n", v, v, v)
	}

	fmt.Println("========================================")

	str_1 := "Hello"
	str_2 := "World"

	str_3 := str_1 + " " + str_2

	fmt.Println(str_3)

	fmt.Println("========================================")

	str_4 := "Hello World!"
	str_5 := str_4

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str_4))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str_5))

	// StringHeader Data field is point to memory address where saved string in memory
	fmt.Println(stringHeader1) // &{4340398307 12}
	fmt.Println(stringHeader2) // &{4340398307 12}

	fmt.Println(unsafe.Sizeof(stringHeader1.Data))

	fmt.Println("========================================")

	var str10 string = "Hello World"
	var slice []byte = []byte(str10)

	slice[2] = 'a' // change

	fmt.Println(str10) // immutable
	fmt.Printf("%s\n", slice)

	fmt.Println("========================================")

	builder := strings.Builder{}

	builder.Write([]byte("Hello"))
	builder.Write([]byte(" World"))

	result := builder.String()

	fmt.Println(result)
}
