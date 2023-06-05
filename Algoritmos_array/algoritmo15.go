package main

import "fmt"

func main() {

	array := [10]float64{1.3, 2.3, 8.7, 1.3, 7.8, 5.9, 17.9, 5.4, 6.1, 10.0}
	var slice []float64
	fmt.Println("Atual lista: ", array)
	for i := 0; i < len(array); i++ {
		if array[i] > 5 {
			slice = append(slice, array[i])
		}
	}
	fmt.Println("Nova lista com números maiores que 5: ", slice)
}
