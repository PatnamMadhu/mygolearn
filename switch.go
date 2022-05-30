package main

import "fmt"

func main() {

x,y := 2,1

//var x,y = 2,1.1
switch x+y {
case 1:
	fmt.Println("Sum value is 1")
case 2:
	fmt.Println("Sum value is 2")
case 3:
	fmt.Println("Sum value is 3")
default:
	fmt.Println("Sum value is default")
}
}
