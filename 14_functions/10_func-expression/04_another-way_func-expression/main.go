package main

import "fmt"

func makeGreeter() func() string {
	return func() string {
		return "Hello fuker"
	}
}

func main() {
	greet := makeGreeter() //func expression, variavel greet, makeGreeter esta sendo assign a variavel greet
	fmt.Println(greet())
}
