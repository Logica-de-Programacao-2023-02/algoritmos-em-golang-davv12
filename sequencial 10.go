package main

import "fmt"

func main() {
	var peso float32
	fmt.Println("qual o seu peso?")
	fmt.Scan(&peso)
	var peso1 float32
	peso1 = peso * 2.20462
	fmt.Println("seu peso em libras é", peso1)

}
