package main

import "fmt"

func main() {
	switch "Jose" {
	case "Arthur":
		fmt.Println("Wassaup Arthur")
	case "Isabela":
		fmt.Println("Wassup Isabela")
	case "Jose":
		fmt.Println("Wassup Jose")
		fallthrough
	case "Hugo":
		fmt.Println("Wassup Hugo")
		fallthrough
	case "Lucas":
		fmt.Println("Wassup Lucas")

	}
}
