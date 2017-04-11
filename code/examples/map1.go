package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var persionDB map[string]PersonInfo
	persionDB = make(map[string]PersonInfo)

	persionDB["12345"] = PersonInfo{"12345", "Tom", "Room 203"}

	persion, ok := persionDB["12345"]

	if ok {
		fmt.Println(persion.Name, persionDB["12345"])
	}
}
