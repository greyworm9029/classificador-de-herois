package main

import (
	"fmt"
)

// Classificador de heróis

func main() {

	// Herói representa as informações de um herói
	type Heroi struct {
		nome string
		exp  int
	}

	heroi := Heroi{
		nome: "Daniel",
		exp:  5000,
	}

	// Condições do programa
	switch {
	case heroi.exp < 1000:
		fmt.Println("Se o XP for menor que 1000 = Ferro")

	case heroi.exp >= 1000 && heroi.exp <= 2000:
		fmt.Println("Se o XP for entre 1001 e 2000 = Bronze")

	case heroi.exp >= 2001 && heroi.exp <= 3000:
		fmt.Println("Se o XP fo entre 2001 e 3000 = Prata")

	case heroi.exp >= 3001 && heroi.exp <= 4000:
		fmt.Println("Se o XP for entre 3001 e 4000 = Ouro")

	case heroi.exp >= 4001 && heroi.exp <= 5000:
		fmt.Println("Se o XP for entre 4001 e 5000 = Platina")

	case heroi.exp >= 5001 && heroi.exp <= 6000:
		fmt.Println("Se o XP for entre 5001 e 6000 = Diamante")

	case heroi.exp >= 6001 && heroi.exp <= 7000:
		fmt.Println("Se o XP for entre 6001 e 7000 = Esmeralda")

	case heroi.exp >= 7001 && heroi.exp <= 8000:
		fmt.Println("Se o XP for entre 7001 e 8000 = Mestre")

	case heroi.exp >= 8001 && heroi.exp <= 9000:
		fmt.Println("Se o XP for entre 8001 e 9000 = Grão Mestre")

	case heroi.exp >= 9001 && heroi.exp <= 10000:
		fmt.Println("Se o XP for entre 9001 e 10000 = Ancião")

	default:
		fmt.Println("Se o XP for igual ou maior do que 10001 = Desafiante")
	}

	for heroi.exp < 5001 {
		fmt.Printf("O jogador de nome %s está no nível %d!\n", heroi.nome, heroi.exp)
		heroi.exp += 5000
	}

}
