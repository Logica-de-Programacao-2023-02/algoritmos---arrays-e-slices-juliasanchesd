package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var x string
	fmt.Print("Digite um valor: ")
	fmt.Scan(&x)

	for i := 0; i < len(slice); i++ {
		if slice[i] == x {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	fmt.Println(slice)
}
