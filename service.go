package main

import (
	"./server"
	"./pb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
	"./core"
)

func main() {
	lis, err := net.Listen("tcp", core.GetConfig().Listen)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	core.InitDB()

	// add new service hanlder here
	pb.RegisterUserServer(s, &server.UserServer{})
	pb.RegisterUserExtraServer(s, &server.UserExtraServer{})

	s.Serve(lis)
}
