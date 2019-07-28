package service

import "m3online/rpc"

var userMap = map[uint32]*rpc.User{}

type UserListService struct {
}

func (service UserListService) GetUserList() (map[uint32]*rpc.User, error) {
	return userMap, nil
}

func (service UserListService) EnterUser(enter *rpc.Enter, user *rpc.User) () {
	userMap[enter.EnterId] = user
}

func (service UserListService) DeleteUser(enter *rpc.Enter) () {

	_, exist := userMap[enter.EnterId]

	if exist {
		delete(userMap, enter.EnterId)
	}
}

func (service UserListService) CheckExist(enter *rpc.Enter) bool {
	_, exist := userMap[enter.EnterId]
	return exist
}
