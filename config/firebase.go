package firebase

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go/kkaycodes"
	"firebase.google.com/go/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	App        *firebase.App
	AuthClient *auth.Client
	MongoClient *mongo.Client
}

func InitializeFirebase(ctx context.Context, credentialsPath string, mongoURI string) (*FirebaseApp, error) {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase Auth: %v", err)
	}

	clientOpts := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	if err := mongoClient.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %v", err)
	}

	return &FirebaseApp{
		App:        app,
		AuthClient: authClient,
		MongoClient: mongoClient,
	}, nil
}

func (fa *FirebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	token, err := fa.AuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	return token, nil
}

func (fa *FirebaseApp) InsertDataToMongo(ctx context.Context, dbName, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := fa.MongoClient.Database(dbName).Collection(collectionName)
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return nil, fmt.Errorf("error inserting document into MongoDB: %v", err)
	}
	return result, nil
}

func (fa *FirebaseApp) Close(ctx context.Context) {
	if err := fa.MongoClient.Disconnect(ctx); err != nil {
		log.Printf("error disconnecting from MongoDB: %v", err)
	}
}
