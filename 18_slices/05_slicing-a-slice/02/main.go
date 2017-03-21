package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2]) // pega comecando do index 1 e só, ou seja index 1
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2]) // pega do inicio até menor que index 2, ou seja index 0 e 1
	fmt.Print("[5:] ")        //pega a partir do index 5 ate o final, ou seja index 5 e 6
	fmt.Println(greeting[5:])
	fmt.Print("[:] ") // pega do index 0 ate final, ou seja, index 0,1,2,3,4,5,6
	fmt.Println(greeting[:])
}
