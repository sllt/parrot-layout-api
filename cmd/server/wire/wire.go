//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sllt/parrot-layout-api/internal/handler"
	"github.com/sllt/parrot-layout-api/internal/repository"
	"github.com/sllt/parrot-layout-api/internal/server"
	"github.com/sllt/parrot-layout-api/internal/service"
	"github.com/sllt/parrot-layout-api/pkg/app"
	"github.com/sllt/parrot-layout-api/pkg/jwt"
	"github.com/sllt/parrot-layout-api/pkg/log"
	"github.com/sllt/parrot-layout-api/pkg/server/http"
	"github.com/sllt/parrot-layout-api/pkg/sid"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	//repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(httpServer *http.Server, job *server.Job) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
