package main

import "fmt"

func main() {
	var linha int
	var coluna int
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("Sua matriz é: ", array)
	fmt.Print("Escolha uma linha. ")
	fmt.Scan(&linha)
	fmt.Print("Escolha uma coluna. ")
	fmt.Scan(&coluna)
	fmt.Println("Posição escolhida", array[linha][coluna])
}
