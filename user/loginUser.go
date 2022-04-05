package user

import (
	"context"
	"os"

	pb "github.com/Bit-Optimizer/Xelo-users/protos"
	"github.com/ISTE-SC-MANIT/megatreopuz-models/utils"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) LoginUser(ctx context.Context,
	req *pb.Empty) (*pb.LoginResponse, error) {
	
		//write login logic here

		decoded, err := utils.GetUserFromFirebase(ctx, s.FirebaseClient)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Could not identify the user")
	}
	database := s.MongoClient.Database(os.Getenv("MONGODB_DATABASE"))
	userCollection := database.Collection(os.Getenv("MONGODB_USERCOLLECTION"))
	count, err := userCollection.CountDocuments(ctx , bson.M{"_id": decoded.UID})
		if err != nil {
		return nil, status.Errorf(codes.Internal, "MongoDB could not check the user entry")
	}
	return &pb.LoginResponse{Status: count!=0},nil
}