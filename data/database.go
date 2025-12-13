package data

import (
	"BFSITGoServer/config"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DataConfig struct {
	MongoURI    string
	MongoClient *mongo.Client
}

//var databseName string = "Schul_Dashboard"
//var studentCollection string = "student"
//var teacherCollection string = "teacher"
//var adminCollection string = "admin"

// Database stellt eine Verbindung zur MongoDB her und speichert den Client in DataConfig.
func Database() {
	uri := config.Config.MongoURI
	// Überprüfen, ob die URI nicht leer ist
	if uri == "" {
		log.Fatal("MongoDB URI konnte nicht gefunden werden in der Konfiguration.")
	}

	// MongoDB-Client verbinden
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Fehler beim Verbinden mit der MongoDB: ", err)
	}

	DataConfig.MongoClient = client
	log.Println("Erfolgreich mit der MongoDB verbunden.")
}
