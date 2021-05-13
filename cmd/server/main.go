package main

import (
	"fmt"
	"github.com/ssuxue/mellow-server/domain/repository"
	"github.com/ssuxue/mellow-server/domain/service"
	"github.com/ssuxue/mellow-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	svc := grpc.NewServer()
	port := 9090

	proto.RegisterUserServiceServer(svc, &service.UserServer{
		Service: repository.NewUserRepository(),
	})

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server start to listen port: %d\n", port)
	reflection.Register(svc)
	if err = svc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
