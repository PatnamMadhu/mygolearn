package main 

import "fmt"

import "mypackage1"

func main(){

x,y := 10,20

//function package_add()

sum := mypackage1.package_add(x,y)
fmt.Println("Sum from package is", sum)

}

