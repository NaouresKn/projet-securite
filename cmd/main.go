//pt d'entrer dans le prog

package main

import (
	"fmt"
	"log"
	"projetSecurite/internal/analyzer"
	"projetSecurite/internal/scanner"
)

func main() {

	fmt.Println("Démarrage du système de sécurité...")

	// Répertoire à scanner
	dir := "./damn.bat"
	files := scanner.ScanBatchFiles(dir)

	fmt.Println("Scanning the file") 

	fmt.Println(files[0])

	// Analyse de chaque fichier trouvé
	for _, file := range files {
		fmt.Println("Analyse du fichier :", file)
		content := scanner.ReadBatchFile(file)

		fmt.Println(content)

		if content == "" {
			log.Println("Impossible de lire le fichier :", file)
			continue
		}

		// Analyse IA
		threatString, threatStatus := analyzer.AnalyzeWithAI(content)


		fmt.Printf("crossed !")



		fmt.Println(threatStatus , threatString)

		// Collecting the malware file
		// executor.HandleFile(file, 0)

		break
	}
}
