//Gestion des configurations (par ex., URL API)
package config

// Config structure les configurations nécessaires
type Config struct {
	APIURL string
}

// LoadConfig retourne les configurations
func LoadConfig() Config {
	return Config{
		APIURL: "https://your-llama-or-gemini-api-url.com/analyze",
	}
}
