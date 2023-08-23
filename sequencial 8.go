package main

import "fmt"

func main() {
	var dia float32
	var diaria float32
	fmt.Println("voce trabalha quantos dias?")
	fmt.Scan(&dia)
	fmt.Println("quanto é a sua diaria?")
	fmt.Scan(&diaria)
	var salario float32
	salario = dia * diaria
	fmt.Println("seu salario é ", salario)
}
