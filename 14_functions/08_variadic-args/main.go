package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // ja tem uma slice, passar a slice
	n := average(data...)                     //data é um iten, e mesmo assim que seja so um,
	//existe um monte de coisa la dentro e estao listados, assim pega esse 1 iten e adiciona ... no final
	fmt.Println(n)
}

func average(sf ...float64) float64 { //esse precisa de argumentos separados por virgula
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
