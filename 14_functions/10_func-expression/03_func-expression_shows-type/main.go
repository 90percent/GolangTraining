package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello World-dd!")
	}
	greeting()
	fmt.Printf("%T \n", greeting)
}
