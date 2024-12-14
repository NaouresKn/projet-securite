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
		threatString, threatStatus := analyzer.AnalyzeWithAI(content , file)

		if threatStatus == 0 {
			fmt.Println("Fichier sain")
		} else {
			fmt.Println("Fichier malveillant")

			// Mise en quarantaine
			executor.HandleFile(file, content)
		}


		
		fmt.Println("Threat status: ", threatStatus)
		fmt.Println("\n")
		fmt.Println("Threat Bilan: ", threatString)

		// Collecting the malware file
		

		break
	}
}
