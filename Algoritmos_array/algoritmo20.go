package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6}

	crescente := true

	for i := 1; i < len(list); i++ {
		anterior := list[i-1]
		atual := list[i]

		if anterior > atual {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("A lista está em ordem crescente")
	} else {
		fmt.Println("A lista está em ordem decrescente")
	}
}
