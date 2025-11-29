package db

import (
	"SwitchTool/schema"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
	"log"
	"fmt"
)

func ConnectToDB() (*mongo.Client , error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    uri := "mongodb://root:netvroot@localhost:27018/?authSource=admin"

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Mongo connect error:", err)
		return nil, err
    }

    // Ping
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Mongo ping error:", err)
		return nil, err
    }

    fmt.Println("Connected to MongoDB!")
    return client,nil
}

func Inventory_Write( client *mongo.Client , inv []schema.Device) {

}