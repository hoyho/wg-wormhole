package rpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/hoyho/wg-wormhole/proto"
	"google.golang.org/grpc"
)

func ServeRpcForever(addr string, vCode string, iface string) error {
	verifyCode = vCode
	registryInterface = iface

	//addr := fmt.Sprintf("%s:%d", network, Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	svc := newRegistryServer()
	pb.RegisterRegistryServer(grpcServer, svc)

	err = svc.preflight()
	if err != nil {
		return err
	}

	fmt.Printf("serving registry grpc at %v \n", addr)
	grpcServer.Serve(lis)
	

	select {}

}
