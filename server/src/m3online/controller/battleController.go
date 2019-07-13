package controller

import (
	"context"
	"io"
	"log"
	"m3online/rpc"
	"m3online/service"
	"time"
)

type BattleController struct {
	EnemyService *service.EnemyService
}

var entryCount = uint32(0)

func (BattleController) EnterBattle(context.Context, *rpc.EnterRequest) (*rpc.EnterResponse, error) {
	entryCount = entryCount + 1
	var enter = &rpc.Enter{EnterId: entryCount}

	log.Printf("%d", entryCount)

	return &rpc.EnterResponse{Enter: enter}, nil
}

func (controller BattleController) Connect(request *rpc.ConnectionRequest, stream rpc.BattleService_ConnectServer) error {
	//ToDO: ConnectionRequestCheck
	for {
		time.Sleep(1 * time.Second)

		if false { // ToDo CheckConnectionBreak
			break
		}
		if false { //ToDo CheckDiff Also Check Dirty
			continue
		}
		enemy, _ := controller.EnemyService.GetEnemy()
		var situation = &rpc.EnemySituation{Enemy: enemy}
		if err := stream.Send(situation); err != nil {
			return err
		}
	}
	return nil
}

func (controller BattleController) Attack(stream rpc.BattleService_AttackServer) error {
	for {
		var attackRequest, err = stream.Recv()
		if err == io.EOF {
			var summery = rpc.SessionSummary{}
			return stream.SendAndClose(&summery)
		}
		if err != nil {
			return err
		}

		//ToDo: AttackRequestCheck
		err = controller.EnemyService.AttackEnemy(attackRequest.Attack)
		if err != nil {
			log.Printf("Error:%s", err)
		}
	}
}

var _ rpc.BattleServiceServer = (*BattleController)(nil)
