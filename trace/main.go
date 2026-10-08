package main

import (
	"fmt"
	"os"
)

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

	err = os.WriteFile(os.Args[2], contenu, 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}
}

