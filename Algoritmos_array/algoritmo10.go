package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	minimo := 999999
	maximo := -99999
	for i := 0; i < len(slice); i++ {
		if minimo > slice[i] {
			minimo = slice[i]
		}
		if maximo < slice[i] {
			maximo = slice[i]
		}
	}
	fmt.Println("Máximo = ", maximo, "Mínimo = ", minimo)
}
