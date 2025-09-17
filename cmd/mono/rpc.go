package main

import (
	idgenv1 "solabar-server/api/idgenerator/v1"
	idgensvc "solabar-server/internal/idgenerator/service"

	"github.com/google/wire"
)

// 初始化本地调用的rpc接口

var RpcProviderSet = wire.NewSet(
	// newUserExtClient,
	newIdGeneratorIntClient,
)

// func newUserExtClient(svc *usersvc.UserExtService) userv1.UserExtClient {
// 	return userv1.NewUserExtClientLocal(svc)
// }

func newIdGeneratorIntClient(svc *idgensvc.IdGeneratorIntService) idgenv1.IdGeneratorIntClient {
	return idgenv1.NewIdGeneratorIntClientLocal(svc)
}
