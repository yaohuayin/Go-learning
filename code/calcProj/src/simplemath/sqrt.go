// 姚华银  Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/*
Package foo implements a set of simple mathematical functions. These comments are for
demonstration purpose only. Nothing more.
If you have any questions, please don’t hesitate to add yourself to
golang-nuts@googlegroups.com.
You can also visit golang.org for full Go documentation.
*/
package simplemath

import "math"

func Sqrt(i int) int {
	v := math.Sqrt(float64(i))
	return int(v)
}
