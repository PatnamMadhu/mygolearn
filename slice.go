package main

import "fmt"

func main() {

	
	arr := [4]int{1,2,3,4}

	fmt.Println("Array:", arr)

	
	slice := arr[:2]

	fmt.Println("Slice:", slice[0])
	
	slice = arr[:3]
	
	fmt.Println("Slice:", slice)
	
	slice = arr[2:3]
	
	fmt.Println(slice, len(slice), cap(slice))

	
}
