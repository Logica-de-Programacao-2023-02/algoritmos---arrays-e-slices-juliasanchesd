package main

import "fmt"

func main() {
	array := [4]float64{1.00, 1.50, 2.00, 2.50}

	produto := 1.0
	for i := 0; i < len(array); i++ {
		produto *= array[i]
	}
	fmt.Printf("O produto dos valores do Array é: %.2f\n", produto)

}
