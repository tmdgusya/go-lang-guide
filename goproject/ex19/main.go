package main

import "fmt"

type account struct {
	balance int
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func (a account) withdrawMethodValueType(amount int) account {
	a.balance -= amount
	return a
}

func main() {
	a := &account{100}

	a.withdrawMethod(20)

	fmt.Printf("Balance : %d \n", a.balance)

	b := a.withdrawMethodValueType(20)

	fmt.Printf("A Balance : %d \n", a.balance)
	fmt.Printf("B Balance : %d \n", b.balance)
}
