package main

import "fmt"

func main() {
	// exercicio declarar uma variavel e mostrar o endereco de memoria dela
	idade := 28
	fmt.Println("Idade: ", idade)
	fmt.Println("Endereço de memoria da Idade: ", &idade)
}
