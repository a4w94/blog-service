package main

import (
	"log"
	"net"
	pb "tag-service/proto"
	"tag-service/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

func main() {
	port = "8001"
	s := grpc.NewServer()

	pb.RegisterTagServiceServer(s, server.NewTagServer())

	//取得RPC方法清單資訊 $grpcurl -plaintext localhost:8000 list
	//取得子類別RPC方法清單資訊 $grpcurl -plaintext localhost:8000 list proto.TagService

	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("s.Serve: %v", err)
	}
}
