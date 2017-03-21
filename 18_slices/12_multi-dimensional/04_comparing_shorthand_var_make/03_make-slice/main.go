package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	//agora que tem length e capacity pode associar ao index. Se nao tiver tem que usar append
	//(como no comentario abaixo)
	student[0] = "Todd"
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
