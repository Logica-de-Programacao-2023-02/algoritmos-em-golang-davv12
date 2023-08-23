package main

import "fmt"

func main() {
	var valor float32
	fmt.Println("qual é o valor do produto?")
	fmt.Scan(&valor)
	var valor1 float32
	valor1 = valor * 0.2
	fmt.Println("o valor do produto com o desconto é", valor1)
}
