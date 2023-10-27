package main

import "fmt"

// init

func init(){
	fmt.Println("init1")
}

func init(){
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
