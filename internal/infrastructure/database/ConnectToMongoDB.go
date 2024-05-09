package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	dbHost := getenv("DB_HOST")
	dbPort := getenv("DB_PORT")
	dbName := getenv("DB_NAME")

	uri := fmt.Sprintf("mongodb://%s:%s/%s", dbHost, dbPort, dbName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, error := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if error != nil {
		log.Fatal(error)
		return nil, error
	}

	return client, nil
}

func getenv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Variável de ambiente não encontrada:", key)
		return ""
	}
	return value
}
