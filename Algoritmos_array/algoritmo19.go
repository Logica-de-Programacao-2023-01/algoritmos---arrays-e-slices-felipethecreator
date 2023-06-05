package main

import "fmt"

func main() {
	var i int
	conj1 := [3]int{5, 20, 42}
	conj2 := [3]int{6, 8, 11}
	sum := [3]int{}

	fmt.Println("Os conjuntos são:")
	fmt.Println(conj1)
	fmt.Println(conj2)

	for i <= 2 {
		sum[i] = conj1[i] + conj2[i]
		i++
	}
	fmt.Println("A soma é: ", sum)
}
