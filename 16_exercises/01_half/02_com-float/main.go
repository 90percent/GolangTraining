package main

import "fmt"

func shrblablous(i int) (float64, bool) {
	return float64(i) / 2, i%2 == 0 //aqui tem que converter o int para float64
}

func main() {
	fmt.Println("Digite um numero inteiro: ")
	var x int

	fmt.Scanf("%d", &x)
	i, b := shrblablous(x)
	fmt.Println("O numero dividido por 2, e se ele é par: ")
	fmt.Println(i, b)
}
