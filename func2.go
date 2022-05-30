package main

import "fmt"

func main() {

	fmt.Println("at the top")

	defer fmt.Println("world")

//defer

	fmt.Println("hello")
	
	defer fmt.Println("at the bottom")
}
