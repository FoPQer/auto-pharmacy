package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var DB *mongo.Database

func RegisterMongodb() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	dbName := os.Getenv("DB_NAME")
	DB = client.Database(dbName)
	fmt.Println("Mongo connected")
}

// Функция для закрытия соединения
func Disconnect() {
	if DB != nil {
		if err := DB.Client().Disconnect(context.Background()); err != nil {
			log.Fatalf("Ошибка при закрытии MongoDB: %v", err)
		}
		fmt.Println("MongoDB отключён.")
	}
}
