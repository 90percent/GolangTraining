package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 1) //length of 0 and capacity of 3, 3 elementos podem ser armazenados na array
	//atualmente a slice esta armazenando nil elements
	// trocar capacidade por outro valor e testar
	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i) // adiciona algo na slice, no caso i
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}
