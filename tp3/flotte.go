package main

import "fmt"

type Appareil struct {
	nom    string
	statut string
}

func main() {
	flotte := []Appareil{
		{"Laptop-01", "critique"},
		{"Serveur-02", "avertissement"},
		{"Imprimante-03", "normal"},
		{"Switch-04", "critique"},
	}

	nouveaux := []Appareil{
		{"Routeur-05", "avertissement"},
		{"NAS-06", "normal"},
	}
	for _, a := range nouveaux {
		flotte = append(flotte, a)
	}

	fmt.Println("=== Rapport de flotte reseau ===")
	fmt.Printf("%d appareils enregistres\n\n", len(flotte))

	for i, appareil := range flotte {
		fmt.Printf("[%d] %s — Statut : %s\n", i+1, appareil.nom, appareil.statut)
		fmt.Println("  Actions declenchees :")

		switch appareil.statut {
		case "critique":
			fmt.Println("  - Alerte critique envoyee au responsable")
			fallthrough
		case "avertissement":
			fmt.Println("  - Ticket d'incident cree")
			fallthrough
		case "normal":
			fmt.Println("  - Entree dans le journal de bord")
		}
		fmt.Println()
	}

	critiques, avertissements := 0, 0
	for i := 0; i < len(flotte); i++ {
		switch flotte[i].statut {
		case "critique":
			critiques++
		case "avertissement":
			avertissements++
		}
	}

	fmt.Println("=== Resume ===")
	fmt.Printf("Critiques    : %d\n", critiques)
	fmt.Printf("Avertissements : %d\n", avertissements)
	fmt.Printf("Normaux      : %d\n", len(flotte)-critiques-avertissements)
}
