package controller

import (
	"context"
	"m3online/rpc"
)

type BattleController struct {
}

func (BattleController) EnterBattle(context.Context, *rpc.EnterRequest) (*rpc.EnterResponse, error) {
	panic("implement me")
}

var _ rpc.BattleServiceServer = (*BattleController)(nil)