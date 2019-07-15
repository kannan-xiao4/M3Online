package controller

import (
	"context"
	"github.com/google/uuid"
	"io"
	"log"
	"m3online/rpc"
	"m3online/service"
	"time"
)

type BattleController struct {
	EnemyService *service.EnemyService
	UserService *service.UserListService
}

var entryCount = uint32(0)

func (controller BattleController) EnterBattle(ctx context.Context, request *rpc.EnterRequest) (*rpc.EnterResponse, error) {
	entryCount = entryCount + 1
	var enter = &rpc.Enter{EnterId: entryCount}
	var user = &rpc.User{} //ToDO: DBから登録済みユーザーを取得
	user.Id = &rpc.UserId{UserId:uuid.New().String()}
	user.UserName = request.UserName

	controller.UserService.EnterUser(enter, user)

	log.Printf("enterId: %d, userName: %s", entryCount, user.UserName)

	return &rpc.EnterResponse{Enter: enter}, nil
}

func (controller BattleController) Connect(request *rpc.ConnectionRequest, stream rpc.BattleService_ConnectServer) error {
	for {
		time.Sleep(1 * time.Second)

		if !controller.UserService.CheckExist(request.Enter) {
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

func (controller BattleController) Exit(ctx context.Context, request *rpc.ExitRequest) (*rpc.ExitResponse, error) {
	controller.UserService.DeleteUser(request.Enter)
	return &rpc.ExitResponse{Result:rpc.ResultCode_SUCCESS}, nil
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

		if !controller.UserService.CheckExist(attackRequest.Attack.Enter) {
			continue
		}

		err = controller.EnemyService.AttackEnemy(attackRequest.Attack)
		if err != nil {
			log.Printf("Error:%s", err)
		}
	}
}

func (controller BattleController) ReceiveUsers(request *rpc.ConnectionRequest, stream rpc.BattleService_ReceiveUsersServer) error {
	for {
		time.Sleep(1 * time.Second)

		if !controller.UserService.CheckExist(request.Enter) {
			break
		}
		if false { //ToDo CheckDiff Also Check Dirty
			continue
		}
		userList, _ := controller.UserService.GetUserList()
		var situation = &rpc.UserList{UserMap:userList}
		if err := stream.Send(situation); err != nil {
			return err
		}
	}
	return nil
}

var _ rpc.BattleServiceServer = (*BattleController)(nil)
