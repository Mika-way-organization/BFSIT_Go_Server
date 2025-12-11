package web

/*
In der Datei handler.go werden die HTTP-Handler-Funktionen definiert, die auf verschiedene Routen reagieren.
Diese Handler verarbeiten eingehende HTTP-Anfragen, führen die erforderliche Logik aus und senden die entsprechenden HTTP-Antworten zurück an den Client.
*/

import (
	"html/template"
	"log"
	"net/http"
)

type IndexData struct {
	Title string
}

// IndexHandler rendert die Index-Seite der Webanwendung.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pagedata := IndexData{
		Title: "Willkommen zu BFSIT Go Server",
	}
	// Template-Datei für die Index-Seite laden und wenn nötig Fehler behandeln
	tmpl, err := template.ParseFiles("web/templates/Index.html")
	if err != nil {
		log.Printf("Fehler beim Laden der Template-Datei: %v", err)
		http.Error(w, "Interner Serverfehler", http.StatusInternalServerError)
		return
	}
	// Template mit den Seitendaten ausführen und an den Client senden
	err = tmpl.Execute(w, pagedata)
	if err != nil {
		log.Printf("Fehler beim Ausführen der Template-Datei: %v", err)
		http.Error(w, "Interner Serverfehler", http.StatusInternalServerError)
	}
}
