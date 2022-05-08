package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hoyho/wg-wormhole/proto"

	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("127.0.0.1:1638", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRegistryClient(conn)

	rep, err := client.GetEndpoint(context.Background(), &pb.GetEndpointRequest{})
	fmt.Printf("rep:%v, err:%v \n", rep, err)
}
