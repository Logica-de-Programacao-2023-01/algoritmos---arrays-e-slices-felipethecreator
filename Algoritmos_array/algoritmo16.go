package main

import "fmt"

func main() {
	array := [10]int{4, 5, 6, 1, 12, 33, 45, 89, 75, 12}
	var result []int

	for _, x := range array {
		if x%2 == 0 {
			result = append(result, x)
		}
	}

	fmt.Println(result)
}
