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
		"Gutten morgen!", //ending comma
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ { //outra maneira de printar os elementos
		fmt.Println(greeting[j])
	}

}
