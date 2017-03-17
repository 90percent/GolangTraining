package main

import "fmt"

func main() {
	m := maior(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(m)
}

func maior(sliceM ...int) int { //funcao maior tem como parametro uma slice.
	//ela recebe uma slice como argumento.
	fmt.Println(sliceM)
	fmt.Printf("%T \n", sliceM)
	var grandao int
	for _, v := range sliceM { // range retorna 2 returns, um é o index e outro valor
		if v > grandao {
			grandao = v
		}
	}
	return grandao
}

//em variadic

//variadic ...parameter
//variadic arg...
