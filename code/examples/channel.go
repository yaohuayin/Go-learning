package main

import (
	"fmt"
	"math/rand"
)

func randGenerator() chan int {
	out := make(chan int)

	go func() {
		i := 0
		for {
			i += 1
			fmt.Println("i = ", i)
			out <- rand.Intn(100)
		}
	}()

	return out
}

func main() {
	randSrvHdl := randGenerator()

	fmt.Printf("%d\n", <-randSrvHdl)
	fmt.Printf("%d\n", <-randSrvHdl)
	fmt.Printf("%d\n", <-randSrvHdl)
	fmt.Printf("%d\n", <-randSrvHdl)
}
