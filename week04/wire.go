//+build wireinject

package main

import (
	"main/internal/biz"
	"main/internal/data"
	"main/internal/service"

	"github.com/google/wire"
)

func InitUserService() *service.UserService {
	wire.Build(service.NewUserService, biz.NewUserBiz, data.NewUserRepo, data.NewDB, data.NewCache)
	return &service.UserService{}
}
