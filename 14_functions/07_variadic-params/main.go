package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64      //acumulador, onde ira juntar as variaveis total := 0.0
	for _, v := range sf { //range vai loop over colection of objects, como é slice ele retorna index e value,
		// como n precisa do index ele coloca _ no lugar de i
		total += v // total = total + v
	}
	return total / float64(len(sf)) // como é um float, divide por um float que é a lenght da slice
}
