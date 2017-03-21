package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	mySlice = append(mySlice[:2], mySlice[3:]...) //enganacao, só tira a posicao que vc nao quer
	//no caso tira a posicao 3, que é o index 2. Assim append index 0, 1 e pula 2, depois acrescenta do index 3
	//(posicao 4) ate o final
	fmt.Println(mySlice)

}
