package main

import "fmt"

func main() {
	switch "Jose" {
	case "Isabela", "Jose":
		fmt.Println("Wassup Isabela, ou, err, Jose")
	case "Hugo", "Lucas":
		fmt.Println("A segunda letra de seus nomes é 's' ")
	case "Arthur", "Lais":
		fmt.Println("Wassup Arthur / Lais")

	}

}
