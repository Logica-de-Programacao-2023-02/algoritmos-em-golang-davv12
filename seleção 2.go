package main

import "fmt"

func main() {
	var numero int
	var numero2 int
	var numero3 int
	fmt.Println("informe o primeiro numero")
	fmt.Scan(&numero)
	fmt.Println("informe o segundo numero")
	fmt.Scan(&numero2)
	fmt.Println("informe o terceiro numero")
	fmt.Scan(&numero3)
	var menor int
	if numero < numero2 {
		menor = numero
	} else {
		menor = numero2
	}
	if numero < numero3 {
		menor = numero
	} else {
		menor = numero3
	}
	if numero2 < numero3 {
		menor = numero2
	} else {
		menor = numero3

	}
	fmt.Println("o menor é", menor)
}
