package service

import (
	"m3online/rpc"
	"time"
)

var currentEnemy = rpc.Enemy{}
var isDirty = false
var connectStream = map[*rpc.Enter]*rpc.BattleService_ConnectServer{}

const (
	masterId   = 0
	enemyMaxHp = 10
	enemyName  = "Slime"
)

type EnemyService struct {
}

func (EnemyService) Initialize() error {
	currentEnemy.InstantId = 0
	currentEnemy.MasterId = masterId
	currentEnemy.Name = enemyName
	currentEnemy.Hp = enemyMaxHp
	return nil
}

func (service EnemyService) Update() error {

	for {
		time.Sleep(1 / 10 * time.Second)

		if isDirty {
			isDirty = false
			situation := &rpc.EnemySituation{Enemy: &currentEnemy}
			for key, stream := range connectStream {
				var err = (*stream).Send(situation)

				if err != nil {
					delete(connectStream, key)
				}
			}
		}
	}

	return nil
}

func getCurrentSituation() rpc.EnemySituation {

}

func (EnemyService) Register(enter *rpc.Enter, stream *rpc.BattleService_ConnectServer) {
	connectStream[enter] = stream
}

func (EnemyService) AttackEnemy(attack *rpc.Attack) error {

	currentEnemy.Hp = currentEnemy.Hp - 1

	if currentEnemy.Hp <= 0 {
		currentEnemy, _ = createEnemy()
	}

	isDirty = true
	return nil
}

func createEnemy() (rpc.Enemy, error) {
	var enemy = rpc.Enemy{}
	enemy.InstantId = currentEnemy.InstantId + 1
	enemy.MasterId = masterId
	enemy.Hp = enemyMaxHp
	enemy.Name = enemyName
	return enemy, nil
}

