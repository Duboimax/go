package main

import (
	"errors"
	"fmt"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("opération inconnue : %s", op)
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		resultat, _ := operer(a, b, op)
		return resultat
	}
}

func main() {
	for {
		var a, b float64
		var op string

		fmt.Print("Entrez deux nombres et une opération (ex: 10 5 +) ou 'quit' pour quitter : ")
		fmt.Scan(&a, &b, &op)

		if op == "quit" {
			fmt.Println("Au revoir !")
			break
		}

		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Printf("Erreur : %s\n", err)
		} else {
			fmt.Printf("%g %s %g = %g\n", a, op, b, resultat)
		}
	}
}
