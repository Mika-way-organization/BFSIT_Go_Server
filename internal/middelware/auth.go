package middelware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func isKeyAllowed(clientKey string, allowedKeys []string) bool {
	if clientKey == "" {
		return false
	}

	for _, key := range allowedKeys {
		if clientKey == key {
			return true
		}
	}
	return false
}

// APIKeyAuthMiddleware überprüft den API-Schlüssel in den Headern der eingehenden Anfragen.
func APIKeyAuthMiddleware(allowedKeys []string) gin.HandlerFunc {
	// Gibt eine Middleware-Funktion zurück
	return func(c *gin.Context) {
		clientKey := c.Query("api_key")
		fmt.Printf("Überprüfe API-Schlüssel: '%s'\n", clientKey)

		// Überprüfen, ob der Client-Schlüssel in der Liste der erlaubten Schlüssel enthalten ist.
		if !isKeyAllowed(clientKey, allowedKeys) {
			// Protokolliere den unautorisierten Zugriff
			log.Printf("Unautorisierter Zugriff auf %s. Übergebener Key: '%s'", c.Request.URL.Path, clientKey)

			// Antwort mit 401 Unauthorized und Abbruch der Anfrage
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "Error",
				"message": "Ungültiger oder fehlender API-Schlüssel.",
			})
			return
		}

		// Wenn der Schlüssel gültig ist, fahre mit der Anfrage fort
		c.Next()
	}
}
