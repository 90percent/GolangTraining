package main

import "fmt"

func main() {
	switch "pumnacara" {

	case "Isabela":
		fmt.Println("Linda!")
	case "Jose":
		fmt.Println("Vc é foda!")
	case "Arthur":
		fmt.Println("Irmao querido")
	default:
		fmt.Println("Vc tem amigos?")

	}
}

// nao precisa de default fallthrough (Exercicio 02)
// fallthrough é opcional
// Voce pode especificar fallthrough ao declarar explicitamente ele
// Break nao é necessario como em outras linguagens
