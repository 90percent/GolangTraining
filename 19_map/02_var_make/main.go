package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	// myGreeting := make(map[string]string) //mesma coisa que a linha de cima, esta no exercicio 3
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
