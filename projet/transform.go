package main

import (
    "strings"
    "unicode"
)

func estMarqueur(mot string) bool {
    return mot == "(up)" || mot == "(low)" || mot == "(cap)"
}

func appliquerMarqueur(mot string, marqueur string) string {
    switch marqueur {
    case "(up)":
        return strings.ToUpper(mot)
    case "(low)":
        return strings.ToLower(mot)
    case "(cap)":
        return capitaliser(mot)
    }
    return mot
}

func capitaliser(mot string) string {
    if mot == "" {
        return mot
    }
    lettres := []rune(strings.ToLower(mot))
    lettres[0] = unicode.ToUpper(lettres[0])
    return string(lettres)
}

func transformerLigneCasse(ligne string) string {
    mots := strings.Fields(ligne)
    motsCorriges := []string{}

    for _, mot := range mots {
        if estMarqueur(mot) {
            if len(motsCorriges) > 0 {
                dernier := motsCorriges[len(motsCorriges)-1]
                motsCorriges[len(motsCorriges)-1] = appliquerMarqueur(dernier, mot)
            }
        } else {
            motsCorriges = append(motsCorriges, mot)
        }
    }

    return strings.Join(motsCorriges, " ")
}
