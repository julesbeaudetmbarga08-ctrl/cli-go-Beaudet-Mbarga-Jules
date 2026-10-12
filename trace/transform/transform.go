// Package transform contient les règles de transformation des logs.
// Toutes ses fonctions sont pures : elles ne lisent ni n'écrivent aucun fichier.
package transform

import (
	"strconv"
	"strings"
	"unicode"
)

// Texte transforme chaque ligne du texte.
func Texte(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		lignes[i] = Ligne(ligne)
	}
	return strings.Join(lignes, "\n")
}

// Ligne transforme une ligne de log : date heure niveau source code lieu.
func Ligne(ligne string) string {
	mots := strings.Fields(ligne)
	if len(mots) != 6 {
		return ligne
	}
	mots[2] = strings.ToUpper(mots[2])
	mots[3] = Capitaliser(mots[3])
	mots[4] = ConvertirCode(mots[4])
	return strings.Join(mots, " ")
}

// Capitaliser met la première lettre en majuscule et le reste en minuscules.
func Capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

// ConvertirCode convertit un code hexadécimal (0x...) ou binaire (0b...) en décimal.
func ConvertirCode(code string) string {
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
