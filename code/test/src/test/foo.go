package main

import (
	"fmt"
)

func main() {
	data := make(chan int)
	exit := make(chan bool)

	data <- 1
	fmt.Println(len(data), cap(data))

	go func() {
		for d := range data {
			fmt.Println(d)
		}

		fmt.Println("recv over.")
		exit <- true
	}()

	data <- 2
	fmt.Println(len(data), cap(data))
	data <- 3
	fmt.Println(len(data), cap(data))
	close(data)

	fmt.Println("send over.")
	<-exit
}
