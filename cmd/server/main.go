package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"BFSITGoServer/internal/api"
	"BFSITGoServer/internal/core/services"
	"BFSITGoServer/web"
)

func main() {
	dataSVC := services.NewTestService() // Instanz des TestService erstellen

	// Router-Konfiguration mit den erforderlichen Diensten
	config := &api.RouterConfig{
		TestService: dataSVC,
	}

	// Gibt mir die Gin-Router-Instanz zurück
	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*.html") //Läd alle HTML Templates
	router.Static("/static", "web/static")      //Statischer Ordner für CSS, JS, Bilder etc.

	router.GET("/", gin.WrapF(web.IndexHandler)) // Index-Route

	apis := router.Group("/api") // Gruppe für Test-API-Routen
	{
		apis.GET("/test", config.GetTestHanlder) // Test-API-Route
	}

	// Starte den Server auf Port 8080
	port := 8080
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Fehler beim Starten des Servers: %v\n", err)
	}

	fmt.Printf("Der Server läuft auf http://127.0.0.1:%d\n", port)
}
