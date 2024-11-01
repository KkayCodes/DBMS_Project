package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
    // Define MongoDB connection URI
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    // Connect to MongoDB
    var err error
    client, err = mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal("Error connecting to MongoDB:", err)
    }
    fmt.Println("Connected to MongoDB!")
}
