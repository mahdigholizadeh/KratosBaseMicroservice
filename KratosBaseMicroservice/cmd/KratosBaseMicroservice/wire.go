//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"KratosBaseMicroservice/internal/biz"
	"KratosBaseMicroservice/internal/conf"
	"KratosBaseMicroservice/internal/data"
	"KratosBaseMicroservice/internal/server"
	"KratosBaseMicroservice/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
