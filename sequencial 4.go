package main

import "fmt"

func main() {
	var idade int
	fmt.Println("qual sua idade?")
	fmt.Scan(&idade)
	var idade2 int
	idade2 = idade * 365
	fmt.Println("sua idade em dias é ", idade2)
}
