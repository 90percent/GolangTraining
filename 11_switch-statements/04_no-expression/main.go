package main

import "fmt"

func main() {

	nomeDoAmigo := "Jose"

	switch {
	case len(nomeDoAmigo) == 2:
		fmt.Println("Wassup my friend with name of lenght 2")
	case nomeDoAmigo == "Jose":
		fmt.Println("Wassup Jose")
	case nomeDoAmigo == "Isabela":
		fmt.Println("Wassup Isabela")
	case nomeDoAmigo == "Arthur":
		fmt.Println("Wassup Arthur")
	case nomeDoAmigo == "Hugo", nomeDoAmigo == "Lucas":
		fmt.Println("Wassup Hugo ou Lucas")
	default:
		fmt.Println("Sem matched, é o default")
	}
}
