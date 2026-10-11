package main

import (
    "os"
)

func lireFichier(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func ecrireFichier(path string, contenu string) error {
    return os.WriteFile(path, []byte(contenu), 0644)
}
