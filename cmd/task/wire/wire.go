//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sllt/parrot-layout-api/internal/server"
	"github.com/sllt/parrot-layout-api/pkg/app"
	"github.com/sllt/parrot-layout-api/pkg/log"
	"github.com/spf13/viper"
)

var serverSet = wire.NewSet(
	server.NewTask,
)

// build App
func newApp(task *server.Task) *app.App {
	return app.NewApp(
		app.WithServer(task),
		app.WithName("demo-task"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		serverSet,
		newApp,
	))
}
