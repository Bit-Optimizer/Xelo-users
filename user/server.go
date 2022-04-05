package user

import (
	"firebase.google.com/go/v4/auth"
	"go.mongodb.org/mongo-driver/mongo"
)

//defining server type for grpc server
type UserServer struct
{
	MongoClient  *mongo.Client
	FirebaseClient *auth.Client
}