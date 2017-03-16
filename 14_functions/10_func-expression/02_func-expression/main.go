package main

import "fmt"

func main() {

	greeting := func()  // colocou a funcao dentro de uma funcao anonima, nao tem nome. e assign para uma variavel,
	//Isso é chamado de func expression, assim vc pode ter uma funcao dentro de uma funcao
		fmt.Println("Hello Worldasss!")
	}

	greeting()
}
