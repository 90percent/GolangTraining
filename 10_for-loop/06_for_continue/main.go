package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		//se o restante for zero, ou seja um numero par, ele n imprime e volta proinicio do for
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		//condicao de parada
		if i >= 50 {
			break
		}
	}
}
