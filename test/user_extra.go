// auto generated file
// DO NOT EDIT!

package main

import (
	"../core"
	"../pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(core.GetConfig().Listen, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserExtraClient(conn)

	r, err := c.GetUserExtra(context.Background(), &pb.UserExtraObject{Uid: 1})

	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Get UserExtra Uid: %d\n", r.Uid)
}
