package bootstrap

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)


func Connect() {
	err := godotenv.Load(".env")
	
	if err != nil{
		log.Fatal(err)
	}


	// app , err := ConnectToFirebase()

	// if err != nil {
	// 	log.Fatalf("error initialising firebase app: %v", err)
	// }

	// client, err := app.Auth(context.Background())
	// if err != nil {
	// 	log.Fatalf("error connecting to firebase user: %v", err)
	// }
	mongo, err := ConnectToMongoDB()

	log.Print(`Pinging MongoDB`)
	log.Print(err)

	err = mongo.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(`Pinged MongoDB successfully`)


}