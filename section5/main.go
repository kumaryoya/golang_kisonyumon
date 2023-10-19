package main

import "fmt"

// 変数

const Pi = 3.14

const (
	URL = "http://xx.co.jp"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// var Big int = 9223372036854775807 + 1
// cannot use 9223372036854775807 + 1 (untyped int constant 9223372036854775808) as int value in variable declaration (overflows)

const Big  = 9223372036854775807 + 1

const (
	c0 = iota
	c1
	c2
)
// 連続する整数の連番

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)
	// cannot assign to Pi (neither addressable nor a map index expression)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)
	// 1 1 1 A A A

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
	// 0 1 2
}
