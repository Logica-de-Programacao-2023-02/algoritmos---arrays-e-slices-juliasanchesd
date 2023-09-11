package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var x int

	fmt.Print("Digite um valor: ")
	fmt.Scan(&x)

	var achou bool

	for _, valor := range array {
		if x == valor {
			achou = true
			break
		}
	}
	if achou {
		fmt.Println("Encontrei o valor de: ", x)
	} else {
		fmt.Println("Não encontrei o valor de: ", x)
	}
}
