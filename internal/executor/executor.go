// Module pour exécuter ou mettre en quarantaine les fichiers
package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// HandleFile exécute ou met en quarantaine un fichier selon son statut
func HandleFile(name string,code string,) { 

	
	fmt.Println("Fichier dangereux détecté ! Mise en quarantaine.")
	quarantineFile(name , code)
	
}

// executeBatchFile exécute un fichier batch
func executeBatchFile(path string) error {
	cmd := exec.Command("cmd", "/C", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// quarantineFile déplace un fichier dans un dossier de quarantaine
func quarantineFile(filePath, content string) {
	// Define the quarantine path
	quarantinePath := "./quarantine/" + filepath.Base(filePath)

	// Ensure the quarantine directory exists
	err := os.MkdirAll("./quarantine/", os.ModePerm)
	if err != nil {
		fmt.Println("Erreur lors de la création du dossier de quarantaine :", err)
		return
	}

	// Create the quarantine file
	file, err := os.Create(quarantinePath)
	if err != nil {
		fmt.Println("Erreur de mise en quarantaine :", err)
		return
	}
	defer file.Close()

	// Write the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erreur de mise en quarantaine :", err)
		return
	}

	fmt.Println("Fichier déplacé en quarantaine :", quarantinePath)
}
