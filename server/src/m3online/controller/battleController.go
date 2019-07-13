package controller

import (
	"context"
	"log"
	"m3online/rpc"
)

type BattleController struct {
}

var entryCount = uint32(0)

func (BattleController) EnterBattle(context.Context, *rpc.EnterRequest) (*rpc.EnterResponse, error) {
	entryCount = entryCount + 1
	var enter = &rpc.Enter{EnterId:entryCount}

	log.Printf("%d", entryCount)

	return &rpc.EnterResponse{Enter:enter}, nil
}

var _ rpc.BattleServiceServer = (*BattleController)(nil)
