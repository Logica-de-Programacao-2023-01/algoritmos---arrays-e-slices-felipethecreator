package main

import "fmt"

func main() {

	var length int

	fmt.Print("Digite o tamanho desejado da slice: ")
	fmt.Scan(&length)

	slice := make([]int, length)

	fmt.Print("\nAgora informe os valores desejados: ")

	for i := 0; i < length; i++ {
		fmt.Scan(&slice[i])
	}

	fmt.Println("\nSeu Slice é: ", slice)

	sum := 0

	for _, result := range slice {
		sum += result
	}

	fmt.Println("\nA soma dos elementos é: ", sum)

}
