package main

import (
	"fmt"
	"math"
)

// Demande à l'utilisateur d'entrer un nombre positif et renvoie cette valeur.
func saisieFloat64Positif(message string) float64 {
	var value float64
	for {
		fmt.Print(message)
		if _, err := fmt.Scan(&value); err == nil && value >= 0 {
			return value
		}
		fmt.Println("Veuillez entrer une valeur positive.")
	}
}

func main() {
	var pDiff, sd, ciLow, ciHigh, ob11, ob12, ob22, ob21, p1, p2, p11, p12, p21, p22, n, confiance, z float64

	// Demander les valeurs des observations
	fmt.Println("**** Calcul de l'intervalle de confiance (niveau 90%, 95%, ou 99%) pour deux proportions appariées ****")
	ob11 = saisieFloat64Positif("Nombre d'observations où les deux résultats sont négatifs (p11) : ")
	ob12 = saisieFloat64Positif("Nombre d'observations où le résultat de la deuxième expérience est positive (p12) : ")
	ob22 = saisieFloat64Positif("Nombre d'observations où les deux résultats sont positifs (p22) : ")
	ob21 = saisieFloat64Positif("Nombre d'observations où la première expérience est positive et la deuxième est négative (p21) : ")

	// Déterminer le niveau de confiance
	for {
		fmt.Print("Veuillez entrer le niveau de confiance (90, 95 ou 99) : ")
		_, err := fmt.Scan(&confiance)
		if err != nil {
			fmt.Println("Veuillez entrer un nombre valide.")
			continue
		}
		if confiance != 90 && confiance != 95 && confiance != 99 {
			fmt.Println("Niveau de confiance invalide. Veuillez entrer 90, 95 ou 99.")
			continue
		}
		break
	}

	// Déterminer la valeur z selon entrée de l'utilisateur
	switch confiance {
	case 90:
		z = 1.645
	case 95:
		z = 1.96
	case 99:
		z = 2.576
	default:
		fmt.Println("Niveau de confiance invalide.")
		return
	}

	// Calcul des proportions
	n = ob11 + ob12 + ob22 + ob21
	p11 = ob11 / n
	p12 = ob12 / n
	p22 = ob22 / n
	p21 = ob21 / n
	p1 = p21 + p22
	p2 = p12 + p22

	// Calcul de la différence des proportions
	pDiff = p1 - p2

	// Calcul de l'écart-type de la différence des proportions
	sd = math.Sqrt((p1*(1-p1) + p2*(1-p2) - 2*((p11*p22)-(p12*p21))) / n)

	// Calcul de l'intervalle de confiance
	ciLow = pDiff - (z * sd)
	ciHigh = pDiff + (z * sd)

	// Affichage du résultat
	fmt.Printf("L'intervalle de confiance à %v%% pour la différence entre les proportions se situe entre [%0.4f, %0.4f]\n", confiance, ciLow, ciHigh)
}
