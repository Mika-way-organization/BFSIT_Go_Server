package services

import (
	"math/rand"
	"time"
)

// TestResponse repräsentiert die Struktur der Testantwort, die vom TestService zurückgegeben wird
type TestResponse struct {
	Message string `json:"message"`
	Joke    string `json:"joke"`
	Time    string `json:"time"`
}

// TestService bietet Methoden zum Abrufen von Testdaten
type TestService struct{}

// NewTestService erstellt eine neue Instanz des TestService
func NewTestService() *TestService {
	return &TestService{}
}

// GetTestData generiert und gibt eine Testantwort zurück
func (ts *TestService) GetTestData() (*TestResponse, error) {
	message := "Hey, das hat ja super geklappt! Hier ist eine Testantwort vom Server... Willst du auch einen Witz hören?"
	jokes := []string{
		"Warum können Programmierer nicht schwimmen? Weil sie nur Quellcode kennen.",
		"Der einzige Ort, wo 'außer Betrieb' steht, ist das Betriebssystem.",
		"Ein Optimist sieht das Glas halb voll. Ein Pessimist sieht das Glas halb leer. Ein Programmierer sieht das Glas doppelt so groß wie nötig.",
		"Warum hat der Entwickler Schluss gemacht? Sie hatte zu viele Schnittstellen und war ihm zu offen (open source).",
		"Was ist das Lieblingsinstrument eines Programmierers? Die C-Harfe.",
		"Was ist der Unterschied zwischen einem Web-Entwickler und einem Web-Designer? Der Web-Designer will, dass es gut aussieht. Der Web-Entwickler will, dass es funktioniert.",
		"Kommt ein Java-Programm in eine Bar. Der Barkeeper sagt: 'Wir bedienen keine Java-Programme.' Darauf das Programm: 'Entschuldigung, dachte, ich hätte hier die VM (Virtual Machine).'",
		"Warum schlafen Server nachts nie? Weil sie Angst haben, dass sie der Systemadministrator im Schlaf reboote (neustartet).",
		"Was macht ein Programmierer, wenn er seekrank wird? Er geht ans Ufer (shore) und schaut, ob er den Fehler im Code beheben kann.",
	}

	//Random Witz auswählen
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]

	currentTime := time.Now().Format(time.RFC1123)

	// TestResponse mit den generierten Daten erstellen
	response := &TestResponse{
		Message: message,
		Joke:    joke,
		Time:    currentTime,
	}

	return response, nil
}
