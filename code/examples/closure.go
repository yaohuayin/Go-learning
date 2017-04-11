package main

import "fmt"

func ExFunc(n int) func() {
	sum := n

	a := func() {
		fmt.Println(sum + 1)
	}

	return a
}

func main() {
	myFunc := ExFunc(10)
	myFunc()

	myAnotherFunc := ExFunc(20)
	myAnotherFunc()

	myFunc()
	myAnotherFunc()
}
