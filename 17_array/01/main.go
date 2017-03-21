package main

import "fmt"

func main() {

	var x [58]int

	fmt.Println(x)      //aqui aparece os zeros pq é uma array de int
	fmt.Println(len(x)) // 58
	fmt.Println(x[42])  // zero pq n tem nada
	x[42] = 777
	fmt.Println(x[42]) // 777 pq colocou ele la
}
