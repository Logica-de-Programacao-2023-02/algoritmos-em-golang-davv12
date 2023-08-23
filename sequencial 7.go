package main

import "fmt"

func main() {
	var numero int
	var numero1 int
	var numero2 int
	fmt.Println("informe um numero")
	fmt.Scan(&numero)
	numero1 = numero + 1
	numero2 = numero - 1
	fmt.Println("o antecessor do numero é", numero2, "o sucessor do numero é", numero1)
}
