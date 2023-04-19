package main

import "fmt"

func main() {

	var a, b, c, d, e, f int
	fmt.Println("Qual o número da linha 0, coluna 0? ")
	fmt.Scan(&a)
	fmt.Println("Qual o número da linha 0, coluna 1? ")
	fmt.Scan(&b)
	fmt.Println("Qual o número da linha 1, coluna 0? ")
	fmt.Scan(&c)
	fmt.Println("Qual o número da linha 1, coluna 1? ")
	fmt.Scan(&d)
	fmt.Println("Qual o número da linha 2, coluna 0? ")
	fmt.Scan(&e)
	fmt.Println("Qual o número da linha 2, coluna 1? ")
	fmt.Scan(&f)

	matriz := [3][2]int{{a, b}, {c, d}, {e, f}}
	fmt.Println("sua matriz:", matriz)
}
