package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}

	soma := 0

	for i := 0; i < len(array); i++ {
		soma += array[i]
	}
	fmt.Printf("A soma dos valores no array é: %d\n", soma)
}
