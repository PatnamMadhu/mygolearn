package main

import "fmt"

func calculator(num1 int , num2 int )(int,int) {

sum := num1 + num2
diff := num1 - num2

return sum, diff

}

func main(){

defer x,y := 20,30

sum, diff := calculator(x,y)

 fmt.Println("Sum",sum)
fmt.Println("Diff",diff)

}
