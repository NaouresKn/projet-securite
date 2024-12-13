//pt d'entrer dans le prog

package main

import (
	"fmt"
	"log"
	"projetSecurite/internal/analyzer"
	"projetSecurite/internal/executor"
	"projetSecurite/internal/scanner"
)

func main() {

	fmt.Println("Démarrage du système de sécurité...")

	// Répertoire à scanner
	dir := "./damn.bat"
	files := scanner.ScanBatchFiles(dir)

	// Analyse de chaque fichier trouvé
	for _, file := range files {
		fmt.Println("Analyse du fichier :", file)
		content := scanner.ReadBatchFile(file)

		if content == "" {
			log.Println("Impossible de lire le fichier :", file)
			continue
		}

		// Analyse IA
		thread := analyzer.AnalyzeWithAI(content)
		// Gestion des résultats
		executor.HandleFile(file, thread)

		break
	}
}
