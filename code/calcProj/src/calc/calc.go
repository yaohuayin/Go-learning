// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/*
Package foo implements a set of simple mathematical functions. These comments are for
demonstration purpose only. Nothing more.
If you have any questions, please don’t hesitate to add yourself to
golang-nuts@googlegroups.com.
You can also visit golang.org for full Go documentation.
*/
package main

import "os" //用于获取命令参数os.Args
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value.")
}

func main() {
	args := os.Args

	//	fmt.Println(len(args), args[0], args[1], args[2], args[3])
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}

		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)

	case "sqrt":
		if len(args) != 2 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}

		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
