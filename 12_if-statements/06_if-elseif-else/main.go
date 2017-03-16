package main

import "fmt"

func main() {

	if false {
		fmt.Println("n imprime pois eh falso, caso alguma coisa: else if :")
	} else if true {
		fmt.Println("imprime pois eh true caso contrario")
	} else {
		fmt.Println("pode ser diferente do primeiro mas n imprime de todo jeito")
	}
}
