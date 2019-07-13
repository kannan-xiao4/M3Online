package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "m3online/rpc"
)

const (
	address = "localhost:10000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBattleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EnterBattle(ctx, &pb.EnterRequest{BattleId:"hogehgoehgoehge"})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("EntryCount: %d", r.Enter.EnterId)

}