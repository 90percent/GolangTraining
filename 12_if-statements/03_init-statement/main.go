package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b { //variavel food fica limitada a essa funcao, scope e talz
		fmt.Println(food)
	}
}
