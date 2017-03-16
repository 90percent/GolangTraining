package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	// so um exemplo contrario do exercicio anterior (03)
	//  tentando acessar a variavel food fora de sua declaracao, scope

	fmt.Println(food)
	// # command-line-arguments
	// ./main.go:15: undefined: food

}
