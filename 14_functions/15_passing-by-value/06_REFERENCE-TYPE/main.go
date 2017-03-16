package main

import "fmt"

func main() {
	m := make([]string, 1, 25) //make slice
	fmt.Println(m)             // [ ]
	changeMe(m)                //passa a slice aqui
	fmt.Println(m)             // [Todd]
}

func changeMe(z []string) { //slice assign to z
	z[0] = "Todd"
	fmt.Println(z) // [Todd]
}
