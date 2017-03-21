package main

import (
	"fmt"
)

func main() {
	var student []string //zero value of slice o que é nil
	var students [][]string
	student[0] = "Todd"
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
