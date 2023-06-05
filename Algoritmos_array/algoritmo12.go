package main

import "fmt"

func main() {

	array := [5]int{1, 3, 6, 5, 12}
	fmt.Println("Array atual:", array)
	slice := []int{}
	for i := 0; i < len(array); i++ {
		if array[i]%3 == 0 {
			slice = append(slice, array[i])
		}
	}
	fmt.Println("Multiplos de 3 dentro do array:", slice)
}
