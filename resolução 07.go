package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}

	var x int

	fmt.Print("Digite um valor: ")
	fmt.Scan(&x)

	var achou bool

	for _, valor := range slice {
		if x == valor {
			achou = true
			break
		}
	}
	if !achou {
		slice = append(slice, x)
	}
	fmt.Println(slice)
}
