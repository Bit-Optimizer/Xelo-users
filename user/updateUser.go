package user

import (
	"context"
	"fmt"
	"os"

	pb "github.com/Bit-Optimizer/Xelo-users/protos"
	"github.com/ISTE-SC-MANIT/megatreopuz-models/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) UpdateUser(ctx context.Context,
	 req *pb.UpdateUserRequest) (*pb.Empty, error){

		fmt.Println("Update User function")
		a := bson.D{
			{"country", req.GetAddress().Country},
			{"division", req.GetAddress().Division},
			{"city", req.GetAddress().City},
			{"lane", req.GetAddress().Lane},
			{"pin", req.GetAddress().Pin},
		}
		fmt.Println(a)
var setUpdateFields bson.D
	if len(req.GetName()) > 0 {
		setUpdateFields = append(setUpdateFields, bson.E{Key: "name", Value: req.GetName()})
	}
	if len(req.GetPhone()) > 0 {
		setUpdateFields = append(setUpdateFields, bson.E{Key: "phone", Value: req.GetPhone()})
	}
	if len(req.GetProfile()) > 0 {
		setUpdateFields = append(setUpdateFields, bson.E{Key: "profile", Value: req.GetProfile()})
	}
	if(len(req.GetAddress().Lane)>0){
		setUpdateFields = append(setUpdateFields, bson.E {Key: "address", Value : a})
	}
	// complete the code here
	decoded , err := utils.GetUserFromFirebase(ctx, s.FirebaseClient)
	if err != nil{
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal , "Internal Server Error")
	}
	database := s.MongoClient.Database(os.Getenv("MONGODB_DATABASE"))
	userCollection := database.Collection(os.Getenv("MONGODB_USERCOLLECTION"))
	res, err := userCollection.UpdateByID(ctx ,
		bson.D{primitive.E{Key: "_id", Value: decoded.UID},},
				bson.D{primitive.E{Key: "$set", Value:setUpdateFields},
		})
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "database refused to create user")
	}
	fmt.Println(res)

		return &pb.Empty{},nil
	 }