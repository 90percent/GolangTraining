package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{} // var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n) //se verdadeiro pega o valor e faz append na variavel xs
		}
	}
	fmt.Printf("%T", callback)
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1 //recebe n como um int e avalia se é > 1 retornando em um bool se é true or false
	})
	fmt.Println(xs) // [2 3 4]
}
