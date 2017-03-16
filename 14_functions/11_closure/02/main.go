package main

import "fmt"

var x = 0 // var x int é mais idiomatico, mais bonito, melhor

func increment() int {
	x++
	return x
}

func main() {

	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope


*/
