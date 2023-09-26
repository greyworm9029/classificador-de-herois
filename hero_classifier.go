package main

import (
	"fmt"
)

// Classificador de heróis

func main() {
	// Heroe armazena o nome e a experiência do herói
	 heroe := struct {
		nome string
		exp int
	} {
		nome: "Rodrigo",
		exp: 1250,
	}

	switch {
	case heroe.exp < 1000:
		fmt.Println("Se o XP for menor do que 1.000 = Ferro");

	case heroe.exp >= 1000 && heroe.exp <= 2000:
		fmt.Println("Se o XP for entre 1.001 e 2.000 = Bronze");

	case heroe.exp >= 2001 && heroe.exp <= 5000:
		fmt.Println("Se o XP for entre 2.001 e 5.000 = Prata");

	case heroe.exp >= 6001 && heroe.exp <= 7000:
		fmt.Println("Se o XP for entre 6.001 e 7.000 = Ouro");

	case heroe.exp >= 7001 && heroe.exp <= 8000:
		fmt.Println("Se o XP for entre 7.001 e 8.000 = Platina");

	case heroe.exp >= 8001 && heroe.exp <= 9000:
		fmt.Println("Se o XP for entre 8.001 e 9.000 = Ascendente");

	case heroe.exp >= 9001 && heroe.exp <= 10000:
		fmt.Println("Se o XP for entre 9.001 e 10.000 = Imortal");

	default: 
	fmt.Println("Se o XP for maior ou igual a 10.001 = Radiante");
	}

	for heroe.exp < 1000 {
		fmt.Printf("O herói com o nome de %s está no nível %d!\n", heroe.nome, heroe.exp)
		heroe.exp += 1000
	}
}
