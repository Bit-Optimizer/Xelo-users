package user

import (
	"context"
	"fmt"

	pb "github.com/Bit-Optimizer/Xelo-users/protos"
)

func (s *UserServer) RegisterUser(ctx context.Context,
	 req *pb.RegisterUserRequest) (*pb.User, error){
		fmt.Println("Register User function")

	// 	u := user.User{
	// 		id:  req.GetId(),
	// 		name:  req.GetName(),
	// 		email:  req.GetEmail(),
	// 		phone:  req.GetPhone(),
	// 		profile:  req.GetProfile(),
	// 		address:  req.GetAddress()
	// 		// rating: 0,
	// 		// ratecount: 0,
	// 		// createdAt: time.Time,
	// 		// updatedAt: time.Time
	// 	}
	// database := s.MongoClient.Database(os.Getenv("MONGODB_DATABASE"))
	// userCollection := database.Collection(os.Getenv("MONGODB_USERCOLLECTION"))
	// nuser, err = userCollection.InsertOne(ctx, u)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, status.Errorf(codes.Internal, "database refused to create user")
	// }
	return &pb.User{
		//code likho idhar
	}, nil
}