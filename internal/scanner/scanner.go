//Module pour scanner les fichiers .bat
package scanner

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// ScanBatchFiles scanne un répertoire à la recherche de fichiers .bat
func ScanBatchFiles(dir string) []string {
	var batchFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(info.Name(), ".bat") {
			batchFiles = append(batchFiles, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Erreur lors du scan :", err)
	}
	return batchFiles
}

// ReadBatchFile lit le contenu d'un fichier .bat
func ReadBatchFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return ""
	}
	return string(content)
}
