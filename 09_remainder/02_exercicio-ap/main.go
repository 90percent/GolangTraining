package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Digite um numero para dividi-lo: ")
	fmt.Scanf("%d", &x)
	fmt.Println("digite o divisor: ")
	fmt.Scanf("%d", &y)
	z := x % y
	fmt.Println("o remainder e: ", z)

}
