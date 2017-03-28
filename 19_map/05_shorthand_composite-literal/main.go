package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	} // valores literais e sao compiste values, pode ter valores

	fmt.Println(myGreeting["Jenny"])
}
