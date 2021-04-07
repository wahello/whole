// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/realotz/whole/internal/conf"
	"github.com/realotz/whole/internal/server"
	"github.com/realotz/whole/internal/services"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
	panic(wire.Build(services.ProviderSet, server.ProviderSet, newApp))
}
