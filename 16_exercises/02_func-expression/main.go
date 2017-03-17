package main

import "fmt"

func main() {
	var x int
	fmt.Println("digite um numero: ")
	fmt.Scanf("%d", &x)
	h := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	// h, b := func(x int) (int, bool) {
	// 	return x / 2, x%2 == 0
	// }
	fmt.Println(h(x)) //
}
