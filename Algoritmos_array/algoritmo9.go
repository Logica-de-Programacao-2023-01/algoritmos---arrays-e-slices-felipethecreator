package main

import "fmt"

func main() {
	var array = [6]float64{}
	var a float64
	fmt.Println("O seu array é ", array, "que número você quer adicionar ao array? ")
	fmt.Scan(&a)

	for i := 0; i < len(array); i++ {
		array[i] = a
	}
	fmt.Println(array)
}
