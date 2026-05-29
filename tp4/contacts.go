package main

import "fmt"

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s, %d ans — %s", p.NomComplet(), p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf(
		"=== Fiche Employe ===\n%s\nPoste   : %s\nAdresse : %s\nSalaire : %.2f EUR",
		e.Presentation(), e.Poste, e.Adresse.Format(), e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pct float64) {
	e.Salaire += e.Salaire * pct / 100
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (et Etudiant) MentionObtenue() string {
	switch {
	case et.Moyenne >= 16:
		return "Tres Bien"
	case et.Moyenne >= 14:
		return "Bien"
	case et.Moyenne >= 12:
		return "Assez Bien"
	case et.Moyenne >= 10:
		return "Passable"
	default:
		return "Insuffisant"
	}
}

func (et Etudiant) FicheEtudiant() string {
	return fmt.Sprintf(
		"=== Fiche Etudiant ===\n%s\nPromo   : %s\nMoyenne : %.2f — Mention : %s",
		et.Presentation(), et.Promo, et.Moyenne, et.MentionObtenue(),
	)
}

func main() {
	e1 := Employe{
		Personne: Personne{"Alice", "Martin", 32, "alice.martin@corp.fr"},
		Adresse:  Adresse{"12 rue de la Paix", "Paris", "75001"},
		Poste:    "Developpeuse Backend",
		Salaire:  3800.00,
	}
	e2 := Employe{
		Personne: Personne{"Bruno", "Leclerc", 45, "b.leclerc@corp.fr"},
		Adresse:  Adresse{"8 avenue Foch", "Lyon", "69006"},
		Poste:    "Chef de projet",
		Salaire:  5200.00,
	}

	e1.AugmenterSalaire(10)
	e2.AugmenterSalaire(5)

	fmt.Println(e1.FicheEmploye())
	fmt.Println()
	fmt.Println(e2.FicheEmploye())
	fmt.Println()

	et1 := Etudiant{
		Personne: Personne{"Clara", "Dupont", 20, "clara.dupont@univ.fr"},
		Promo:    "BUT Info 2025",
		Moyenne:  17.5,
	}
	et2 := Etudiant{
		Personne: Personne{"Dylan", "Morel", 21, "dylan.morel@univ.fr"},
		Promo:    "BUT Info 2025",
		Moyenne:  11.8,
	}

	fmt.Println(et1.FicheEtudiant())
	fmt.Println()
	fmt.Println(et2.FicheEtudiant())
}
