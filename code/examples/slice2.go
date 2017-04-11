package main

import "fmt"

func main() {
	mySlice1 := []int{1, 2, 3, 4, 5}
	mySlice2 := []int{5, 4, 3}

	fmt.Println("mySlice1: ", mySlice1)

	copy(mySlice1, mySlice2)
	fmt.Println("mySlice1: ", mySlice1)

	fmt.Println("mySlice2: ", mySlice2)

}
