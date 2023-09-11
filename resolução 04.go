package main

import "fmt"

func main() {
	var tamanho, valor int

	fmt.Print("Digite o tamanho da Slice: ")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um valor: ")
		fmt.Scan(&valor)
		slice[i] = valor
	}

	var soma int

	for _, x := range slice {
		soma += x
	}
	fmt.Println("Meu slice é: ", slice)
	fmt.Println("A soma é: ", soma)
}
