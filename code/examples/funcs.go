package main

func main() {
	var x interface{}
	//var a int = 2

	x = nil

	switch i := x.(type) {
	case nil:
		println("x is nil")

	case int:
		println(i)
	}
}
