package main

import "fmt"

func main() {
	var x [256]string
	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		x[i] = string(i) //pegando o numero e transformando em uma letra e vendo como o slice of bytes é
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))

	}
}
