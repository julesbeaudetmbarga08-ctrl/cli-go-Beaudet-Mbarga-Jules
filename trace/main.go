package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// transformerTexte transforme chaque ligne du texte.
func transformerTexte(texte string) string {
	lignes := strings.Split(texte, "\n")

	for i, ligne := range lignes {
		lignes[i] = transformerLigne(ligne)
	}

	return strings.Join(lignes, "\n")
}

// transformerLigne transforme une ligne de log :
// date heure niveau source code lieu
func transformerLigne(ligne string) string {
	mots := strings.Fields(ligne)

	if len(mots) != 6 {
		return ligne
	}

	mots[2] = strings.ToUpper(mots[2])
	mots[3] = capitaliser(mots[3])
	mots[4] = convertirCode(mots[4])

	return strings.Join(mots, " ")
}

// capitaliser met la première lettre en majuscule et le reste en minuscules.
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}

	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])

	return string(lettres)
}

// convertirCode convertit un code hexadécimal (0x...) ou binaire (0b...) en décimal.
func convertirCode(code string) string {
	base := 0
	if strings.HasPrefix(code, "0x") {
		base = 16
	} else if strings.HasPrefix(code, "0b") {
		base = 2
	} else {
		return code
	}
	nombre, err := strconv.ParseInt(code[2:], base, 64)
	if err != nil {
		return code
	}
	return strconv.FormatInt(nombre, 10)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage : go run . <entree> <sortie>")
		return
	}

	contenu, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	resultat := transformerTexte(string(contenu))

	err = os.WriteFile(os.Args[2], []byte(resultat), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}

	fmt.Println("Transformation terminée :", os.Args[2])
}
