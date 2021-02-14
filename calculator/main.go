package main

import (
	"os"
	"os/signal"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	"github.com/snathale/challenge-hash/calculator/application"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	GenerateConfigFileError = errors.New("impossible generate example config")
)

func main() {
	var fileName string
	var config string
	appCli := &cli.App{
		Name:  "calculate discount",
		Usage: "grpc server that calculate product discount",
	}
	appCli.Commands = []*cli.Command{
		{
			Name:    "config example generator",
			Aliases: []string{"csg"},
			Action: func(c *cli.Context) error {
				return generateConfigSample(fileName)
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "file, f",
					Value:       "./config.json",
					Usage:       "config file name",
					Destination: &fileName,
				},
			},
		},
		{
			Name:    "run application",
			Aliases: []string{"run"},
			Action: func(c *cli.Context) error {
				return run()
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "config, c",
					Value:       "./config.json",
					Usage:       "config application file",
					Destination: &config,
				},
			},
		},
	}
	if err := appCli.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func generateConfigSample(fileName string) error {
	if err := application.NewConfigFile(fileName); err != nil {
		return errors.WithMessage(err, "impossible generate example config")
	}
	return nil
}

func run() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	var err error
	if err = viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
	config := &application.Config{}
	if err = viper.Unmarshal(config); err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(logrus.Level(config.LogLevel))
	var app *application.Application
	if app, err = application.NewApp(config); err != nil {
		return err
	}
	errChan := app.Run()
	return waitForGracefullyShutdown(app, errChan)
}

func waitForGracefullyShutdown(app *application.Application, errChan <-chan error) error {
	interruptServiceChan := make(chan os.Signal)
	signal.Notify(interruptServiceChan, os.Interrupt)
	defer app.Close()
	select {
	case err := <-errChan:
		return err
	case <-interruptServiceChan:
		logrus.Warning("gracefully shutdown")
	}
	return nil
}
