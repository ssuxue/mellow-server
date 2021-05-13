package main

import (
	"context"
	"github.com/ssuxue/mellow-server/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	const address = "localhost:9090"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	cli := proto.NewUserServiceClient(conn)
	req := &proto.LoginRequest{
		Email: "chase",
		Password: "123456",
	}

	res, err := cli.Login(context.Background(), req)
	if err != nil {
		log.Println(err)
	}

	log.Printf("id of the user is: %s", res.Id)
}
