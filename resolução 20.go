package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	cres := true
	for i := 0; i < len(array)-1; i++ {
		atl := array[i]
		prox := array[i+1]
		if atl >= prox {
			cres = false
			break
		}
	}

	if cres {
		fmt.Print("Arry está ordenado.")
	} else {
		fmt.Print("Não está ordenado.")
	}
}
