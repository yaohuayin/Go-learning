package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Binary int64

func (i Binary) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Binary) Get() int64 {
	return int64(i)
}

func main() {
	b := Binary(200)
	s := Stringer(b)
	fmt.Println(s.String())
}
