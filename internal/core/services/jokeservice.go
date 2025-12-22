package services

import (
	"math/rand"
	"time"
)

type JokeResponse struct {
	Joke string `json:"joke"`
	Time string `json:"time"`
}

type JokeService struct{}

func NewJokeService() *JokeService {
	return &JokeService{}
}

func (js *JokeService) GetJoke() (*JokeResponse, error) {
	jokes := []string{
		"Oma: \"Das blöde Waffeleisen funktioniert nicht!\" Enkel: \"Geh sofort weg von meinem Laptop!\"",
		"Kellner: \"Wollen sie ihren Kaffee schwarz?\" Gast: \"Was für Farben haben sie noch?\"",
		"Lehrer: \"Geh bitte vor die Tür Peter! Dein Gerede interessiert keinen!\" Schüler: \"Dann können sie gleich mitkommen!\"",
		"Geht ein Dalmatiner einkaufen. Fragt ihn die Kassiererin: \"Sammeln sie noch Punkte?\"",
		"Ich habe eben bei \"Weight Watchers\" angerufen. Es nimmt keiner ab.",
		"Wo geht eine Wal essen? Im Wahllokal.",
		"Was macht ein Pirat am Computer? Er drückt die Enter-Taste.",
		"Warum können Geister so schlecht lügen? Weil man durch sie hindurchsehen kann.",
		"Was ist orange und läuft durch den Wald? Eine Wanderine.",
		"Treffen sich zwei Jäger. Beide tot.",
		"Was steht auf einem Grabstein von einem IT-Experten? 'Endlich offline.'",
		"Was steht auf dem Grabstein eines Programmierers? 'Hier liegt ein Code, der nie gelaufen ist.'",
		"Warum sind Programmierer schlechte Boxer? Weil sie immer in der Schleife hängen bleiben.",
		"Was steht auf dem Grabstein eines Diabetikers? - Sein Leben war kein Zuckerschlecken.",
		"Was bestellt ein Idiot beim Italiener? Trottelini.",
		"Ein Medizin-Professor fragt: \"Nennen Sie mir ein eisenhaltiges Abführmittel.\" Student: \"Eine Handschelle.\"",
	}

	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]

	currentTime := time.Now().Format(time.RFC1123)

	response := &JokeResponse{
		Joke: joke,
		Time: currentTime,
	}
	return response, nil
}
