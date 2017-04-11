package main

//import "fmt"

const (
	a, b = iota, iota << 10
	c, d
)

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {
	println(c)
}

func main() {
	c := Black

	test(c)S

	//x := 1

	test(1)
}

%HOMEDRIVE%%HOMEPATH%
