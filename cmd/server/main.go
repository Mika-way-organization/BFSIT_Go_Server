package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"BFSITGoServer/config"
	"BFSITGoServer/data"
	"BFSITGoServer/internal/api"
	"BFSITGoServer/internal/core/services"
	"BFSITGoServer/internal/middelware"
	"BFSITGoServer/web"
)

var ExpectedAPIKeys = []string{}

func init() {
	config.LoadConfig() // Lädt alle Werte und speichert sie in config.Config

	// JETZT, da config.Config gefüllt ist, können wir die Keys für die Middleware vorbereiten.
	ExpectedAPIKeys = []string{
		config.Config.GruppeOneAPIKey,
		config.Config.GruppeTwoAPIKey,
		config.Config.GruppeThreeAPIKey,
		config.Config.GruppeFourAPIKey,
		config.Config.GruppeFiveAPIKey,
	}
}

func main() {
	//Datenbank verbindung Testen
	data.Database()

	dataSVC := services.NewTestService() // Instanz des TestService erstellen
	jokeSVC := services.NewJokeService() // Instanz des JokeService erstellen

	// Router-Konfiguration mit den erforderlichen Diensten
	config := &api.RouterConfig{
		TestService: dataSVC,
		JokeService: jokeSVC,
	}

	// Gibt mir die Gin-Router-Instanz zurück
	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*.html") //Läd alle HTML Templates
	router.Static("/static", "web/static")      //Statischer Ordner für CSS, JS, Bilder etc.

	router.GET("/", gin.WrapF(web.IndexHandler)) // Index-Route

	apis := router.Group("/api") // Gruppe für Test-API-Routen
	{
		apis.Use(middelware.APIKeyAuthMiddleware(ExpectedAPIKeys)) // API-Key-Middleware anwenden
		apis.GET("/test", config.GetTestHanlder)                   // Test-API-Route
		apis.GET("/joke", config.GetJokeHandler)                   // Witz-API-Route
	}

	// Starte den Server auf Port 8080
	port := 8080
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Der Server startet auf http://127.0.0.1%s\n", addr)

	err := router.Run(addr)

	// Dieser Teil wird nur ausgeführt, wenn router.Run() mit einem Fehler beendet wird.
	if err != nil {
		// gin.Run verwendet intern log.Fatal, aber wenn Sie den Fehler hier selbst behandeln wollen,
		fmt.Printf("Kritischer Fehler beim Beenden des Servers: %v\n", err)
	}
}
