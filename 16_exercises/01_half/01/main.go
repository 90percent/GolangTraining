package main

import "fmt"

func shrblablous(i int) (int, bool) {
	return i / 2, i%2 == 0
}

func main() {
	fmt.Println("Digite um numero inteiro: ")
	var x int

	fmt.Scanf("%d", &x)
	i, b := shrblablous(x) // se a funcao retorna dois valores, deve ter duas variaveis para receber cada um
	fmt.Println("O numero dividido por 2, e se ele é par: ")
	fmt.Println(i, b)
}
