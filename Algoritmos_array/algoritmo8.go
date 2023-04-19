package main

import "fmt"

func main() {

	var nome string
	lista := []string{"Pericles", "Cazuza", "Gabigol", "Diesel", "Caçote", "Vanessa", "Mikael", "Beto"}
	fmt.Println("Nomes: ", lista, "Quem você gostaria de tirar? ")
	fmt.Scan(&nome)

	for i := 0; i < len(lista); i++ {
		if nome == lista[i] {
			fmt.Println("Sua nova lista:", append(lista[:i], lista[i+1:]...))
		}
	}
}
