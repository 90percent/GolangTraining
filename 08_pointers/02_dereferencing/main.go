package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

}

// b é do tipo int pointer
// b aponta para o endereço de memoria onde um int é armazenado
// para ver o valor armazenado neste endereço, adicionamos um * na frente do b
// isto é conhecido como desreferenciar
// o * asterisco é um operador nesse caso
