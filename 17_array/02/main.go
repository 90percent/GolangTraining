package main

import "fmt"

func main() {
	var x [58]string
	// tem de ser 58 apesar de 122 - 65 = 57 , a diferenca é menor mas esta armazenando todos entao 58
	fmt.Println(x) //aqui aparece uma array vazia pq o tipo é string e nao aparece os zeros
	fmt.Println(len(x))
	fmt.Println(x[42]) //acessando um iten na posicao 42, que é o elemento 43 por causa do index 0

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i) //conversao de i em uma string, 65 letra A, na posicao de index 0 e vai crescendo
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
