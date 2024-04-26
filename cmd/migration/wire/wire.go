//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sllt/parrot-layout-api/internal/repository"
	"github.com/sllt/parrot-layout-api/internal/server"
	"github.com/sllt/parrot-layout-api/pkg/app"
	"github.com/sllt/parrot-layout-api/pkg/log"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewUserRepository,
)
var serverSet = wire.NewSet(
	server.NewMigrate,
)

// build App
func newApp(migrate *server.Migrate) *app.App {
	return app.NewApp(
		app.WithServer(migrate),
		app.WithName("demo-migrate"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serverSet,
		newApp,
	))
}
