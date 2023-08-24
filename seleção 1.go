package main

import "fmt"

func main() {
	var numero int
	var numero2 int
	fmt.Println("informe um numero")
	fmt.Scan(&numero)
	fmt.Println("informe um numero")
	fmt.Scan(&numero2)
	var maior int
	if numero > numero2 {
		maior = numero
	} else {
		maior = numero2
	}
	fmt.Println("o maior é", maior)
}
