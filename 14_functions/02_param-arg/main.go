package main

import "fmt"

func main() {
	greet("Jane") // A chamada da funcao greet recebe o ARGUMENT Jane
	greet("John")
}

func greet(name string) { // a funcao greet tem como PARAMETER o name
	fmt.Println(name)
}

// greet is declared with a parameter
// when calling greet, pass in an argument
