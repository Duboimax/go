package main

import (
	"fmt"
	"math"
	"strings"
)

// ══ Interface Payeur — satisfaite implicitement par tout mode de paiement ══
type Payeur interface {
	Payer(montant float64) (string, error)
}

// ══ CarteCredit ══
type CarteCredit struct {
	Numero    string
	Titulaire string
	Solde     float64
}

func (cc *CarteCredit) Payer(montant float64) (string, error) {
	if montant > cc.Solde {
		return "", fmt.Errorf("solde insuffisant : %.2f€ disponible pour %.2f€ demandé", cc.Solde, montant)
	}
	cc.Solde -= montant
	quatreDerniers := cc.Numero
	if len(cc.Numero) >= 4 {
		quatreDerniers = cc.Numero[len(cc.Numero)-4:]
	}
	return fmt.Sprintf("Transaction CB #%s confirmée", quatreDerniers), nil
}

// ══ PayPal ══
type PayPal struct {
	Email string
	Solde float64
}

func (pp *PayPal) Payer(montant float64) (string, error) {
	if montant > pp.Solde {
		return "", fmt.Errorf("solde insuffisant : %.2f€ disponible pour %.2f€ demandé", pp.Solde, montant)
	}
	pp.Solde -= montant
	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, pp.Email), nil
}

// ══ Crypto ══
const tauxBTC = 50000.0 // 1 BTC = 50000€

type Crypto struct {
	Adresse string
	Solde   float64
	Monnaie string
}

func (c *Crypto) Payer(montant float64) (string, error) {
	if montant > c.Solde {
		return "", fmt.Errorf("solde insuffisant : %.2f€ disponible pour %.2f€ demandé", c.Solde, montant)
	}
	c.Solde -= montant
	quantite := math.Round(montant/tauxBTC*1000) / 1000
	return fmt.Sprintf("Paiement de %.3f %s (%.2f€) vers %s", quantite, c.Monnaie, montant, c.Adresse), nil
}

// Vérifications statiques à la compilation
var _ Payeur = &CarteCredit{}
var _ Payeur = &PayPal{}
var _ Payeur = &Crypto{}

// ══ ProcesserPanier — accepte n'importe quel Payeur ══
func ProcesserPanier(payeur Payeur, articles []float64) error {
	total := 0.0
	for _, prix := range articles {
		total += prix
	}
	fmt.Printf("Total du panier : %.2f€\n", total)

	switch payeur.(type) {
	case *CarteCredit:
		fmt.Println("Mode de paiement : Carte de crédit")
	case *PayPal:
		fmt.Println("Mode de paiement : PayPal")
	case *Crypto:
		fmt.Println("Mode de paiement : Crypto-monnaie")
	default:
		fmt.Println("Mode de paiement : inconnu")
	}

	confirmation, err := payeur.Payer(total)
	if err != nil {
		return err
	}
	fmt.Println(confirmation)
	return nil
}

func main() {
	articles := []float64{29.99, 15.50, 120.00}

	cb := &CarteCredit{Numero: "4532015112830366", Titulaire: "Alice Martin", Solde: 500}
	if err := ProcesserPanier(cb, articles); err != nil {
		fmt.Println("Erreur :", err)
	}
	fmt.Println(strings.Repeat("-", 40))

	pp := &PayPal{Email: "bruno.leclerc@mail.com", Solde: 200}
	if err := ProcesserPanier(pp, articles); err != nil {
		fmt.Println("Erreur :", err)
	}
	fmt.Println(strings.Repeat("-", 40))

	crypto := &Crypto{Adresse: "1A1zP1eP5QGefi2DMPTfTL5SLmv7Divf", Solde: 1000, Monnaie: "BTC"}
	if err := ProcesserPanier(crypto, articles); err != nil {
		fmt.Println("Erreur :", err)
	}
	fmt.Println(strings.Repeat("-", 40))

	pauvre := &CarteCredit{Numero: "4532015112830366", Titulaire: "Clara Dupont", Solde: 50}
	if err := ProcesserPanier(pauvre, articles); err != nil {
		fmt.Println("Erreur :", err)
	}
}
