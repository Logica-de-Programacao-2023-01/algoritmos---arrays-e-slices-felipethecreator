package main

import "fmt"

func main() {

	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var num int
	fmt.Print("Informe um número que possa estar contido na array: ")
	fmt.Scan(&num)

	contido := false

	for i := 0; i < len(array); i++ {
		if array[i] == num {
			contido = true
		}
	}
	if contido {
		fmt.Println(num, "pertence a array.")
	} else {
		fmt.Println(num, "não pertence a array.")
	}
}
