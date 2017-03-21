package main

import "fmt"

func main() {

	mySlice := []int{1, 3, 5, 7, 9, 11} //nao tem numero entre os [ ] , entao é um slice, uma slice of int ou ints
	fmt.Printf("%T \n", mySlice)        //[]int
	fmt.Println(mySlice)                // [1 3 5 7 9 11]
}
