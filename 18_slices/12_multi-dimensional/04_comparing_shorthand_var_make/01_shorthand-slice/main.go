package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Todd" //erro pq nao foi setado a length do underlying array
	// student = append(student, "Todd") //isso funciona comentando o de cima /\
	fmt.Println(student)
	fmt.Println(students)
}
