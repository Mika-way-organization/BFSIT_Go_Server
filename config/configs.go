package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config ist die zentrale Struktur, die alle geladenen Werte enthält.
var Config struct {
	MongoURI          string
	WeatherAPIKey     string
	EmailPassword     string
	GruppeOneAPIKey   string
	GruppeTwoAPIKey   string
	GruppeThreeAPIKey string
	GruppeFourAPIKey  string
	GruppeFiveAPIKey  string
}

// Liste der kritischen Keys für die Validierung
var requiredKeys = []string{
	"MONGO_URI", "WEATHER_API_KEY", "EMAIL_PASSWORD",
	"GRUPPE_ONE_API_KEY", "GRUPPE_TWO_API_KEY", "GRUPPE_THREE_API_KEY", "GRUPPE_FOUR_API_KEY", "GRUPPE_FIVE_API_KEY",
}

// LoadConfig lädt die .env-Datei und validiert alle kritischen Keys.
func LoadConfig() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Printf("WARNUNG: .env-Datei nicht gefunden. Versuche Keys aus Systemumgebung zu laden: %v", err)
	}

	// Validierung und Zuweisung der kritischen Keys
	for _, key := range requiredKeys {
		value, found := os.LookupEnv(key)

		// Prüfung auf Existenz und Nicht-Leere
		if !found || value == "" {
			log.Fatalf("Kritischer Fehler: Umgebungsvariable '%s' ist nicht gesetzt oder leer. Server kann nicht starten.", key)
		}

		// Zuweisung zu der exportierten Config-Struktur
		switch key {
		case "MONGO_URI":
			Config.MongoURI = value
		case "WEATHER_API_KEY":
			Config.WeatherAPIKey = value
		case "EMAIL_PASSWORD":
			Config.EmailPassword = value
		case "GRUPPE_ONE_API_KEY":
			Config.GruppeOneAPIKey = value
		case "GRUPPE_TWO_API_KEY":
			Config.GruppeTwoAPIKey = value
		case "GRUPPE_THREE_API_KEY":
			Config.GruppeThreeAPIKey = value
		case "GRUPPE_FOUR_API_KEY":
			Config.GruppeFourAPIKey = value
		case "GRUPPE_FIVE_API_KEY":
			Config.GruppeFiveAPIKey = value
		}
	}
	log.Println("Alle kritischen Umgebungsvariablen erfolgreich geladen und geprüft.")
}
