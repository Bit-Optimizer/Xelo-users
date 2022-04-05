package user

import (
	"context"
	"fmt"
	"os"

	"github.com/Bit-Optimizer/Xelo-models/user"
	pb "github.com/Bit-Optimizer/Xelo-users/protos"
	"github.com/ISTE-SC-MANIT/megatreopuz-models/utils"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) GetUser(ctx context.Context,
	req *pb.Empty) (*pb.User, error) {
	
		//write User Query
		decoded , err := utils.GetUserFromFirebase(ctx, s.FirebaseClient)
	if err != nil{
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal , "Internal Server Error")
	}
	database := s.MongoClient.Database(os.Getenv("MONGODB_DATABASE"))
	userCollection := database.Collection(os.Getenv("MONGODB_USERCOLLECTION"))
	var u = user.User{}
 	 err = userCollection.FindOne(ctx , bson.M{"_id": decoded.UID} ).Decode(&u)
	  if err != nil{
		  return nil, status.Errorf(codes.Internal , "Cannot Find the User")  
	  }
	  
	return &pb.User{},nil
}