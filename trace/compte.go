package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func compterJetons() {
	f, _ := os.Open("tests/T4-entree.txt") // ou T5
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ligne := scanner.Text()
		jetons := strings.Fields(ligne)
		fmt.Printf("Ligne : %q → %d jetons\n", ligne, len(jetons))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur de lecture : %v\n", err)
	}
}
