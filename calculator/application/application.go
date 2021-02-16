package application

import (
	"github.com/sirupsen/logrus"
	"github.com/snathale/challenge-hash/calculator/application/controller"
	"github.com/snathale/challenge-hash/calculator/infrastucture"
	"github.com/snathale/challenge-hash/calculator/interface/server"
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
		return nil, err
	}
	ctrl := controller.NewController(rep)
	if grpcServer, err = server.NewServer(config.Server, ctrl); err != nil {
		return nil, err
	}
	return &Application{
		grpcServer: *grpcServer,
		controller: ctrl,
	}, nil
}

func (a *Application) Run() <-chan error {
	logrus.Info("app run")
	return a.grpcServer.Run()
}

func (a *Application) Close() {
	a.grpcServer.Close()
}
