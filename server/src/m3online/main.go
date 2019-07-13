package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"m3online/service"
	"net"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	controller "m3online/controller"
	pb "m3online/rpc"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 10000, "The server port")
)

func serialize(enter *pb.Enter) string {
	return fmt.Sprintf("%d", enter.EnterId)
}

func newServer() *controller.BattleController {
	var es = &service.EnemyService{}
	es.Initialize()

	s := &controller.BattleController{EnemyService: es}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBattleServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
