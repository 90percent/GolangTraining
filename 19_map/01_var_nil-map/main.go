package main

import "fmt"

//Descomentar a func main para trocar entre os modos de declaracao

func main() {

	var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}

// add these lines:
/*					[key]  =  "value"
myGreeting["Tim"] = "Good morning."
myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map
//
//
// func main() {
//
// 	var myGreeting = map[string]string{}
// 	myGreeting["Tim"] = "Good morning."
// 	fmt.Println(myGreeting)
// 	fmt.Println(myGreeting == nil)
// }
