package application

import (
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	"github.com/snathale/challenge-hash/calculator/application/controller"
	"github.com/snathale/challenge-hash/calculator/infrastucture"
	"github.com/snathale/challenge-hash/calculator/interface/server"
)

var (
	ApplicationGrpcNewServerError = errors.New("impossible create grpc server")
	ApplicationRepositoryError    = errors.New("impossible create repository")
	ApplicationControllerError    = errors.New("impossible create controller")
)

type Application struct {
	grpcServer server.Server
	controller controller.Controller
}

func NewApp(config *Config) (*Application, error) {
	var grpcServer *server.Server
	var err error
	var rep *infrastucture.Repository
	if rep, err = infrastucture.NewRepositories(config.Db); err != nil {
		logrus.WithError(err).Warning(ApplicationRepositoryError)
		return nil, ApplicationRepositoryError
	}
	ctrl := controller.NewController(rep)
	if grpcServer, err = server.NewServer(config.Server, *ctrl); err != nil {
		logrus.WithError(err).Warning(ApplicationGrpcNewServerError)
		return nil, ApplicationGrpcNewServerError
	}
	return &Application{
		grpcServer: *grpcServer,
		controller: *ctrl,
	}, nil
}

func (a *Application) Run() <-chan error {
	logrus.Info("app run")
	return a.grpcServer.Run()
}

func (a *Application) Close() {
	a.grpcServer.Close()
}
