// Module pour exécuter ou mettre en quarantaine les fichiers
package executor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// HandleFile exécute ou met en quarantaine un fichier selon son statut
func HandleFile(path string, thread int) { // path est le chemin a traiter
	if thread == 0 {
		fmt.Println("Fichier sûr détecté. Exécution :", path)
		err := executeBatchFile(path)
		if err != nil {
			fmt.Println("Erreur d'exécution :", err)
		}
	} else {
		fmt.Println("Fichier dangereux détecté ! Mise en quarantaine.")
		quarantineFile(path)
	}
}

// executeBatchFile exécute un fichier batch
func executeBatchFile(path string) error {
	cmd := exec.Command("cmd", "/C", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// quarantineFile déplace un fichier dans un dossier de quarantaine
func quarantineFile(filePath string) {
	quarantinePath := "./quarantine/" + filepath.Base(filePath)
	err := os.Rename(filePath, quarantinePath)
	if err != nil {
		fmt.Println("Erreur de mise en quarantaine :", err)
		return
	}
	fmt.Println("Fichier déplacé en quarantaine :", quarantinePath)
}
