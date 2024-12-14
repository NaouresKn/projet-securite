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
	"strings"

	"github.com/go-gomail/gomail"
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






func SendMail(target string , fileName string, status string , bilan string , code string) error {
	
	loadEnvError := godotenv.Load()

	if loadEnvError != nil {	
		fmt.Printf("Error loading the env file : " , loadEnvError)
		return loadEnvError
	}

	google := "tmakaveli643@gmail.com"
	pass := "svbj ozuu bfzo pqxq"

	if google == "" || pass == "" {
		fmt.Printf("Error loading the google email and password")
		return loadEnvError
	}

	m := gomail.NewMessage()
	m.SetHeader("From", google)
	m.SetHeader("To", target)
	m.SetHeader("Subject", "New Scan Report TDM WITH KNANI NAOURES X ALAA BARKA")
	m.SetBody("text/html", fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	  <head>
		<style>
		  body {
			font-family: Arial, sans-serif;
			background-color: #f9f9f9;
			margin: 0;
			padding: 0;
		  }
		  .container {
			max-width: 600px;
			margin: 30px auto;
			background-color: #ffffff;
			padding: 20px;
			border-radius: 10px;
			box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
		  }
		  h1 {
			color: #333;
			font-size: 22px;
		  }
		  p {
			color: #555;
			line-height: 1.6;
			font-size: 16px;
		  }
		  .code {
			margin: 20px 0;
			padding: 15px;
			background-color: #f4f4f4;
			border-left: 4px solid #007BFF;
			font-size: 20px;
			font-weight: bold;
			color: #333;
		  }
		  .footer {
			margin-top: 20px;
			font-size: 14px;
			color: #888;
		  }
		</style>
	  </head>
	  <body>
		<div class="container">
		  <h1>Status of your scanned file : `+fileName+`</h1>
		  <p>With Status = `+status+`</p>

		  <h3>
		 	Bilan : `+bilan+` 
		  </h3>


		  <p>
		 	Here's the Code Source to verify : 
			`+code+` 
		 
		 
		  </p>

		  <p>We will collect the file to be quarantined..</p>
		  <br/>
		  <div class="footer">
			&copy; 2024 Barka Alaa X Knani Naoures Corporation. All rights reserved.
		  </div>
		</div>
	  </body>
	</html>
	`))

	d := gomail.NewDialer("smtp.gmail.com", 587, google, pass)

	if err := d.DialAndSend(m); err != nil	{
		log.Println(err)
		return err
	}
	return nil
}


// AnalyzeWithAI envoie le contenu du fichier à une API IA
func AnalyzeWithAI(content string, fileName string) (string , int ) {

	loadEnvError := godotenv.Load()
	if loadEnvError != nil {
		fmt.Printf("Error loading the env file : " , loadEnvError)
		return "" , -1
	}


	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		fmt.Printf("Error loading the api key of the gemini model")
	}


	firstPayload := "You will analyse an echo bat file script and you will check if the file is safe or not like does it contain threat or not, is it a virus or a malware or not" 

 
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
					bilan : should be 1 line explaination  *&* as separaters between the status and the bilan  
					status : 1 if threat, otherwise not 
				`},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON payload: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	candidates, ok := responseMap["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		fmt.Println("No candidates found in the response")
		return "", -1
	}
	firstCandidate := candidates[0].(map[string]interface{})
	contentMap := firstCandidate["content"].(map[string]interface{})
	parts := contentMap["parts"].([]interface{})
	text := parts[0].(map[string]interface{})["text"].(string)
	responseParts := strings.Split(text , "*&*")
	if len(responseParts) != 2 {
		fmt.Println("Error in the response format")
		return "" , -1
	}
	status := responseParts[0]
	bilan := responseParts[1]
	statusId := strings.Split(status, ":")[1]
	statusId = strings.TrimSpace(statusId)
	statusIdNumber := 0
	if statusId == "1" {
		statusIdNumber = 1
	}
	fmt.Println("Status : " , statusIdNumber)



	//client who will receive the report 
	clientMail := "knaninaoures5@gmail.com"
	
	errSending := SendMail(clientMail , fileName , status , bilan , content)
	if errSending != nil {
		fmt.Println("Error sending the email")
	}


	return bilan , statusIdNumber	
}
