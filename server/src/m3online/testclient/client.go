package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	enterBattle(c, ctx)

	go connectBattle(c, ctx)

	attackEnemy(c, ctx)
}

func enterBattle(client pb.BattleServiceClient, ctx context.Context) {
	r, err := client.EnterBattle(ctx, &pb.EnterRequest{BattleId: "hogehgoehgoehge"})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("EntryCount: %d", r.Enter.EnterId)
}

func connectBattle(client pb.BattleServiceClient, ctx context.Context) {
	stream, err := client.Connect(ctx, &pb.ConnectionRequest{})

	if err != nil {
		log.Fatalf("%v.ConnectBattle(_) = _, %v", client, err)
	}
	for {
		enemySituation, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.EnemySituation(_) = _, %v", client, err)
		}
		log.Println(enemySituation)
	}
}

func attackEnemy(client pb.BattleServiceClient, ctx context.Context) {

	stream, err := client.Attack(ctx)
	if err != nil {
		log.Fatalf("%v.Attack(_) = _, %v", client, err)
	}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		_ = stdin.Text()

		var request = &pb.AttackRequest{}
		if err := stream.Send(request); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, request, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
}
