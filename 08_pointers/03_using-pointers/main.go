package main

import "fmt"

func main() {

	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42 //o valor nesse endereco de memoria vai para 42
	fmt.Println(a)

}

// isso é importante
// nos podemos passar o endereco de memoria ao invez de um monte de valores
// (nos passamos uma REFERENCIA)
// isso faz nossos programas mais eficiente
// nos nao precisamos passar grandes quantidades de dados
// nos passamos apenas o endereco
//
// tudo é transmitido por valor em go
// quando passamos endereco de memoria, nos estamos passando (transmitindo) um valor.
