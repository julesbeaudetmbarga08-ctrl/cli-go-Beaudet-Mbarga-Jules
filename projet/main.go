package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage : correcteur <input> <output>")
        return
    }

    input := os.Args[1]
    output := os.Args[2]

    contenu, err := lireFichier(input)
    if err != nil {
        fmt.Println("Erreur lecture :", err)
        return
    }

    lignes := strings.Split(contenu, "\n")
    resultat := []string{}

    for _, ligne := range lignes {
        resultat = append(resultat, transformerLigneCasse(ligne))
    }

    err = ecrireFichier(output, strings.Join(resultat, "\n"))
    if err != nil {
        fmt.Println("Erreur écriture :", err)
        return
    }

    fmt.Println("Transformation terminée :", output)
}
