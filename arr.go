package main

import "fmt"

func main() {
	var numeros [5]int
	numeros = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

}
