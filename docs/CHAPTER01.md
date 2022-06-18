# Go Lang 

## 폴더 생성

```
goproject/hello/extra
```

위와 같은 폴더구조를 지니게 됬다면, hello 폴더 아래있는 모든 `.go` 파일들은 hello package 에 포함됩니다. 그리고 만약 
extra package 에 포함되어 있는 `.go` 파일들은 extra package 에 포함되게 됩니다. 즉, `go` 에서는 폴더가 달라지면 package 도 달라집니다.
Kotlin 과 Java 와는 약간 다른 구조다.

## create Go module

```sh
go mod init goproject/hello
```

go module 에는 NPM 진영의 package.json 과 같이 모듈명과 go-version 및 필요한 정보가 담겨있습니다.

## Go build

```sh
go build
```

위 명령어를 입력하면 go file 들이 build 되며, Architecture 나 OS 따라서도 다르게 빌드할 수 있습니다.

```sh
GOOS=linux GOARCH=arm64 go build
```

## Code 분석

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

위의 Code 를 보면 package 가 main 임을 알 수 있는데, 이는 **Go 에서 특별한 의미를 가진 패키지로 프로그램 진입점을 알려주는 패키지**입니다. 
또한 Go 언어는 반드시 package 선언이 먼져 이루어져야 합니다.

### import "fmt"

Go 에서 package 를 import 해오는 방법입니다.

### main func

Go Lang 또한 main 함수가 시작 진입점 입니다.

## Go lang 변수 선언

Go lang 에서 변수를 선언 하는 방법은 아래와 같습니다.

```
변수 선언 키워드 | 변수 이름 | 타입 = 초기값
```

그래서 위의 코드를 아래 실제 코드로 나타내면 아래와 같습니다.

```go
func main() {
	var minimumWage int = 10
	var workingHour int = 10

	var income int = minimumWage * workingHour

	fmt.Println(minimumWage, workingHour, income)
}
```

### Go 변수명 규칙

Go lang 에서 변수를 지을때는 반드시 문자나 _ 로 시작해야 합니다.

- 올바른 예 : abc, _roach
- 올바르지 않은 예 : 1abc, 123

그리고 Go 는 기본적으로 camelCase 를 권장합니다.

### Go Variable Size

이름 | 설명 | 값의 범위 |
--- | --- | --- |
uint8 | 1바이트 없는 부호 정수 | 0 ~ 255 |
uint16 | 2바이트 없는 부호 정수 | 0 ~ 65535 |
uint32 | 4바이트 없는 부호 정수 | 0 ~ 4294967295 |
uint64 | 8바이트 없는 부호 정수 | 0 ~ 18446744073709551615 |
int8 | 1바이트 있는 부호 정수 | -128 ~ 127 |
int16 | 2바이트 있는 부호 정수 | -32768 ~ 32767 |
int32 | 4바이트 있는 부호 정수 | -2,147,483,648 ~ 2,147,483,647 |
int64 | 8바이트 있는 부호 정수 | -9,223,372,036,854,775,808 ~ -9,223,372,036,854,775,809|
float32 | 4바이트 실수 | IEEE-754 32 비트 실수
float64 | 4바이트 실수 | IEEE-754 64 비트 실수
complex64 | 8바이트 복소수(진수, 가수) | 진수와 가수 범위는 float32 와 같음
complex128 | 16바이트 복소수(진수, 가수) | 진수와 가수 범위는 float64 와 같음
byte | uint8 의 별칭 | 0 ~ 255 |
rune | int32 의 별칭, UTF-8 문자 하나를 나타낼때 사용 | -2,147,483,648 ~ 2,147,483,647 |
int | architecture 마다 int32, int64 | - |
uint | architecture 마다 uint32, uint64 | - |

### 다양한 변수 선언 방법


```go
func main() {
	var a int = 10
	var b int
	c := 7 // var 키워드와 Type 생략

	fmt.Println(a, b, c)
}
```

### 서로 다른 타입간 연산

```go
func main() {
	a := 3
	var b float64 = 3.5

	var c int = a
	// d = a * b 서로 다른 타입간 연산 불가

	var e int64 = 7
	// f := a * e 같은 int 이지만 서로 크기가 다르므로 연산 불가

	fmt.Println(a, b, c, e)
}
```

위 처럼 서로 다른 타입일 경우 연산이 불가능하다.

```go
func main() {
	a := 3
	var b float64 = 3.5

	var c int = a
	d := float64(a * c)

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
```

위와 같이 형변환을 해서 연산하는 경우 연산이 가능하다.