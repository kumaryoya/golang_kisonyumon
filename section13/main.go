package main

import (
	"fmt" //推奨
	f"fmt"
	."fmt"
	"example.com/section13/foo"
)

// スコープ

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "BBB"
		fmt.Println(b)
	}
	fmt.Println(b)
	return b
}

func main() {
	fmt.Println(foo.Max)
	// fmt.Println(foo.min)
	f.Println(foo.ReturnMin())
	Println(foo.ReturnMin())

	fmt.Println(appName())
	// fmt.Println(AppName)
	// fmt.Println(Version)

	fmt.Println(Do("AAA"))
}
