package main

import "fmt"

func main() {
	var peso int
	var altura int
	fmt.Println("qual é seu peso?")
	fmt.Scan(&peso)
	fmt.Println("qual é sua altura?")
	fmt.Scan(&altura)
	var altura1 int
	altura1 = altura * altura
	fmt.Println("o seu IMC é ", peso/altura1)
}
