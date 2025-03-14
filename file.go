package main

import (
	"log"
	"os"
)

func removeFileAndDir(paths []string) {
	for _, path := range paths {
		err := os.RemoveAll(path)

		if err != nil {
			log.Printf("Erreur lors de la suppression du fichier/dossier %s. %s", path, err)
		}
	}
}
