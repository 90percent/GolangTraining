package main

//package scope

import "fmt"

func main() {
	x := 42
	//scope de x vai ate o final de main
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring"
		fmt.Println(y)
	}

	//fmt.Println(y) // nao deve ficar aqui pois sai do scope de y

}
