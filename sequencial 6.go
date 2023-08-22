package main

import "fmt"

func main() {
	var salario float32
	fmt.Println("qual seu salario?")
	fmt.Scan(&salario)
	var salario1 float32
	var salario2 float32
	salario2 = 0.15
	salario1 = ((salario * salario2) + salario)
	fmt.Println("seu salario é ", salario1)

}
