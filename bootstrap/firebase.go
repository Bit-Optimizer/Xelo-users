package bootstrap

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func ConnectToFirebase()(*firebase.App , error) {
	
	opt := option.WithCredentialsFile("m2.json")
	app, err := firebase.NewApp(context.Background(), nil,opt)
	if err != nil {
		return nil, fmt.Errorf("error in connectiong firebase %v",err)
	}

	return app, nil
}