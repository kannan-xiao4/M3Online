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

	var enter = enterBattle(c, ctx)

	go connectBattle(c, ctx, enter)
	go getUserList(c, ctx, enter)

	attackEnemy(c, ctx, enter)
}

func enterBattle(client pb.BattleServiceClient, ctx context.Context) *pb.Enter {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()

	r, err := client.EnterBattle(ctx, &pb.EnterRequest{BattleId: "hogehgoehgoehge", UserName: stdin.Text()})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("EntryCount: %d", r.Enter.EnterId)

	return r.Enter
}

func connectBattle(client pb.BattleServiceClient, ctx context.Context, enter *pb.Enter) {
	stream, err := client.Connect(ctx, &pb.ConnectionRequest{Enter: enter})

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

func getUserList(client pb.BattleServiceClient, ctx context.Context, enter *pb.Enter) {
	stream, err := client.ReceiveUsers(ctx, &pb.ConnectionRequest{Enter: enter})

	if err != nil {
		log.Fatalf("%v.ReceiveUsers(_) = _, %v", client, err)
	}
	for {
		userList, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.UserList(_) = _, %v", client, err)
		}
		log.Println(userList)
	}
}

func attackEnemy(client pb.BattleServiceClient, ctx context.Context, enter *pb.Enter) {

	stream, err := client.Attack(ctx)
	if err != nil {
		log.Fatalf("%v.Attack(_) = _, %v", client, err)
	}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		var text = stdin.Text()

		if text == "quit" {
			break
		}

		var request = &pb.AttackRequest{Attack: &pb.Attack{Enter: enter}}
		if err := stream.Send(request); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, request, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)

	r, err := client.Exit(ctx, &pb.ExitRequest{Enter: enter})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	log.Printf("ExistResult: %s", r.Result.String())
}
