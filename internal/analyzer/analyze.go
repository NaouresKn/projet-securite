// Module pour appeler l'IA Gemini/Llama
package analyzer

import (
	"bytes" // pour la manipulation des donnees en memoire
	"encoding/json"
	"fmt" // gerrer les E/Spour afficher messages d'erreurs
	"net/http"
)

// AnalysisRequest pour envoyer une requete à l'IA
type AnalysisRequest struct {
	Content string `json:"content"`
} // Content contient le texte a analyser

// AnalysisResponse la reponse de l IA
// si 0 alors thread sinon 1
type AnalysisResponse struct {
	Thread int    `json:"thread"`
	Reason string `json:"reason"`
}

// AnalyzeWithAI envoie le contenu du fichier à une API IA
func AnalyzeWithAI(content string) int {
	apiURL := "https://"
	requestBody := AnalysisRequest{Content: content}
	// instance qui contient le contenu a analyser

	jsonData, err := json.Marshal(requestBody) //Convertit l'objet requestBody en JSON pour qu'il puisse être envoyé dans la requête HTTP
	if err != nil {
		fmt.Println("Erreur de préparation de la requête :", err)
		return 1
	}

	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erreur lors de l'appel à l'API :", err)
		return 1
	}
	defer resp.Body.Close()

	var result AnalysisResponse
	err = json.NewDecoder(resp.Body).Decode(&result) //Décode la réponse
	if err != nil {
		fmt.Println("Erreur lors du décodage de la réponse :", err)
		return 1
	}

	return result.Thread
}
