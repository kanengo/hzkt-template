//go:build wireinject
// +build wireinject

package main

import (
	"log/slog"
	userbiz "solabar-server/internal/user/biz"

	userservice "solabar-server/internal/user/service"

	idgenbiz "solabar-server/internal/idgenerator/biz"
	idgenser "solabar-server/internal/idgenerator/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func wireApp(*slog.Logger) (*kratos.App, error) {
	panic(wire.Build(
		userbiz.ProviderSet, userservice.ProviderSet,
		idgenser.ProviderSet, idgenbiz.ProviderSet,
		RpcProviderSet,
		newServer, newApp,
	))
}
