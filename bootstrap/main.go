package bootstrap

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/Bit-Optimizer/Xelo-users/protos"
	"github.com/Bit-Optimizer/Xelo-users/user"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)


func Connect() {

	//Loading Env Variables
	err := godotenv.Load(".env")	
	if err != nil{
		log.Fatal(err)
	}

	//Establishing Firebase Connection
	app , err := ConnectToFirebase()

	if err != nil {
		log.Fatalf("error initialising firebase app: %v", err)
	}

	//Creating Firebase Client
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error connecting to firebase user: %v", err)
	}

	//Establishing MongoDB conncection
	mongo, err := ConnectToMongoDB()

	log.Print(`Pinging MongoDB`)
	log.Print(err)

	//Checking Connection
	err = mongo.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(`Pinged MongoDB successfully`)
	
	//Abstracting Port
	port := os.Getenv("PORT")

	//Listining to port
	lis,err := net.Listen("tcp",port)

	if err != nil{
		log.Fatalf("Error Listining in port %v , %v" ,port, err)
	}

	//Adding Additional data to struct
	userServer := user.UserServer{
		MongoClient: mongo,
		FirebaseClient:  client,
	}

	log.Printf("Server listing to port %v" ,port)
	log.Printf("starting Grpc Server")

	// Creating Grpc Server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServer)


	//Listing server req to Port 
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatalf("Could not start gRPC server. Error: %v", err.Error())
	}
}