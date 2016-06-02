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

	c := pb.NewUserClient(conn)

	r, err := c.GetUser(context.Background(), &pb.UserObject{Uid: 1})

	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Get User Uid: %d\n", r.Uid)
}
