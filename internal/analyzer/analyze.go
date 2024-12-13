// Module pour appeler l'IA Gemini/Llama
package analyzer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// pour la manipulation des donnees en memoire

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
func AnalyzeWithAI(content string) (string , int ) {




	
	//process of loading the api key 



	loadEnvError := godotenv.Load()
	if loadEnvError != nil {
		fmt.Printf("Error loading the env file : " , loadEnvError)
		return "" , -1
	}




	//loading the api key 

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		fmt.Printf("Error loading the api key of the gemini model")
	}







	//process of sending the payload to the gemini model 



	firstPayload := "You will analyse an echo bat file script and you will check if the file is safe or not like does it contain threat or not, is it a virus or a malware or not" 








	//http client : 
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey

	// Define the JSON payload
	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": `
					`+firstPayload+`

					here's the code 
					
					`+content+`
					

					Do not forget to return results other than this format : 


					bilan : should be 1 line explaination 
					status : 1 if threat, otherwise not 
				`},
				},
			},
		},
	}

	// Marshal the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON payload: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)






	return "" , 0
}
