//go:build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"lottery/internal/conf"
	"lottery/internal/data"
	"lottery/internal/server"
	"lottery/internal/service"
	"lottery/internal/task"
	"lottery/internal/usecase"
)

func initApp(*conf.Server, *conf.Data) (*kratos.App, error) {
	panic(wire.Build(DataProviderSet, ServiceProviderSet, UsecaseProviderSet, SeverProviderSet, newApp))
}

var DataProviderSet = wire.NewSet(
	data.NewData,
	data.NewDatabase,
	data.NewCache,
	data.NewCouponRepo,
	data.NewPrizeRepo,
	data.NewResultRepo,
	data.NewBlackIpRepo,
	data.NewBlackUserRepo,
	data.NewLotteryTimesRepo,
)

var ServiceProviderSet = wire.NewSet(
	service.NewLotteryService,
	service.NewAdminService,
)

var UsecaseProviderSet = wire.NewSet(
	usecase.NewAdminCase,
	usecase.NewLotteryCase,
	usecase.NewLimitCase,
)

var SeverProviderSet = wire.NewSet(
	task.NewTaskServer,
	server.NewGRPCServer,
	server.NewHTTPServer,
)
