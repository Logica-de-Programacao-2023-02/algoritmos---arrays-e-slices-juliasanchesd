package main

import "fmt"

func main() {
	var matriz [3][2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Digite o valor da linha %d e coluna %d: ", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}
	for _, ints := range matriz {
		fmt.Println(ints)
	}

}
