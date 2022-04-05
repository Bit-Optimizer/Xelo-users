package user

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Bit-Optimizer/Xelo-models/user"
	pb "github.com/Bit-Optimizer/Xelo-users/protos"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) RegisterUser(ctx context.Context,
	 req *pb.RegisterUserRequest) (*pb.Empty, error){
		fmt.Println("Register User function")
		a := user.Address{
			Country: req.GetAddress().Country,
			Division: req.GetAddress().Division,
			City: req.GetAddress().City,
			Lane: req.GetAddress().Lane,
			Pin: req.GetAddress().Pin,
		}
		u := user.User{
			Id:  req.GetId(),
			Name:  req.GetName(),
			Email:  req.GetEmail(),
			Phone: req.GetPhone(),
			Profile:  req.GetProfile(),
			Address: a,
			Rating: 0,
			Ratecount: 0,
			Wishlist: []string{},
			Products: []string{},
			CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
			UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		}

	database := s.MongoClient.Database(os.Getenv("MONGODB_DATABASE"))
	userCollection := database.Collection(os.Getenv("MONGODB_USERCOLLECTION"))
	res, err := userCollection.InsertOne(ctx, u)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "database refused to create user")
	}
	fmt.Println(res)
	return &pb.Empty{}, nil
}