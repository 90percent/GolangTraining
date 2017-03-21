package main

import "fmt"

func main() {
	mySlice := make([]int, 1) //length and cap to 1, underlying array tb se for acrescentar algo tem q ser com append
	fmt.Println(mySlice[0])   // 0
	mySlice[0] = 7            // se for index 1 n daria certo, tem q ser com append
	//mySlice = append(mySlice, 8)
	fmt.Println(mySlice[0]) // 7
	mySlice[0]++            // mySlice[0] = mySlice[0] + 1
	// adiciona 1 ao valor que esta no index 0, entao fica 7 + 1 = 8
	// ou qualquer maneira de adicao, += 1
	fmt.Println(mySlice[0]) //8
}
