package main

import "fmt"

type emp struct{

name string
address string
age int
//i string

}

func display_emp(e emp){
fmt.Println(e.name)
}

func main(){

var empdetail1 emp

empdetail1.name="Dibya"
empdetail1.address="Bng"
empdetail1.age=100

empdetail2 := emp{"Sumit", "Delhi", 200}

display_emp(empdetail1)
display_emp(empdetail2)

}
