package bootstrap

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongoDB() (*mongo.Client,error) {
	connectionString := "mongodb+srv://bit:optimizer@cluster0.6ukqp.mongodb.net/xelo?retryWrites=true&w=majority"
	clientOption := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)

	defer cancel()

	client,err := mongo.Connect(ctx,clientOption)

	if err != nil{
		return nil, fmt.Errorf("error connecting to mongodb: %v", err)
	}

	return client,nil
}