package main

import (
	"fmt"

	"example.com/section14/section14/alib"
)

// テスト

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))

	s := []int{1,2,3,4,5}
	fmt.Println(alib.Average(s))
}
