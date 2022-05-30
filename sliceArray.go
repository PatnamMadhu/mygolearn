package main
import "fmt"

func main(){

x := [6] string {"Name1","Name2","Name3","Name4"}

fmt.Println("Array after initialization, no slice", x)

//Creating the first slice
var y [] string = x[1:4]

//Creating the second slice
var z [] string = x[0:2]

fmt.Println("After slice creation", y)

fmt.Println("After creating the second slice creation", z)

//how to modify values in the slice; Name2 --> Brillio2

y[1]="Brillio2"

fmt.Println("After modifying the slice", y)

fmt.Println("After modifying the slice printing original array X", x)

//length of the slice
fmt.Println("Length of Slice Y",len(y));

//append 

z = append(y,z...)

fmt.Println("Appended Slices Y&Z together",z)

}
