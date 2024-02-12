package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Arturyus92/chat-server/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User created: %+v", req.String())

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("User updated: %+v", req.GetText())

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("User deleted: %+v", req.String())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
