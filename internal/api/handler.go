package api

import (
	"BFSITGoServer/internal/core/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouterConfig hält die Abhängigkeiten für die API-Handler
type RouterConfig struct {
	TestService *services.TestService
	JokeService *services.JokeService
}

// GetTestHanlder behandelt die Anfragen an die Test-API und gibt eine Testantwort zurück
func (rc *RouterConfig) GetTestHanlder(router *gin.Context) {
	data, err := rc.TestService.GetTestData()
	if err != nil {
		log.Printf("Fehler beim Abrufen der Testdaten: %v", err)
		router.JSON(http.StatusInternalServerError, gin.H{"error": "Interner Serverfehler"})
		return
	}
	router.JSON(http.StatusOK, gin.H{"data": data})
}

func (rc *RouterConfig) GetJokeHandler(router *gin.Context) {
	data, err := rc.JokeService.GetJoke()
	if err != nil {
		log.Printf("Fehler beim Abrufen des Witzes: %v", err)
		router.JSON(http.StatusInternalServerError, gin.H{"error": "Interner Serverfehler"})
		return
	}
	router.JSON(http.StatusOK, gin.H{"data": data})
}
