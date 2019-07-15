package service

import "m3online/rpc"

var currentEnemy = rpc.Enemy{}

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

func (EnemyService) GetEnemy() (*rpc.Enemy, error) {
	return &currentEnemy, nil
}

func createEnemy() (rpc.Enemy, error) {
	var enemy = rpc.Enemy{}
	enemy.InstantId = currentEnemy.InstantId + 1
	enemy.MasterId = masterId
	enemy.Hp = enemyMaxHp
	enemy.Name = enemyName
	return enemy, nil
}

func (EnemyService) AttackEnemy(attack *rpc.Attack) error {

	currentEnemy.Hp = currentEnemy.Hp - 1

	if currentEnemy.Hp <= 0 {
		currentEnemy, _ = createEnemy()
	}

	return nil
}
