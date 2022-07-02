package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	
	_ "github.com/micheltank/eth-fee-calculator/docs"
	"github.com/micheltank/eth-fee-calculator/internal/infra/config"
	"github.com/micheltank/eth-fee-calculator/internal/port/rest"
)

func main() {
	appConfig, err := config.NewConfig()
	if err != nil {
		logrus.WithError(err).Fatalf("failed to read config")
	}
	setLogLevel(appConfig)

	err = run(appConfig)
	if err != nil {
		logrus.WithError(err).Fatal("failed running application")

		return
	}
}

func setLogLevel(appConfig config.Config) {
	logLevel, err := logrus.ParseLevel(appConfig.LogLevel)
	if err != nil {
		logrus.WithError(err).Warnf("failed to parse log level, using default 'info'")

		return
	}
	logrus.SetLevel(logLevel)
}

func run(appConfig config.Config) error {
	logrus.Info("Starting application")

	// REST Server
	restApiServer, err := rest.NewServer(appConfig)
	if err != nil {
		return errors.Wrap(err, "failed to initialize restApiServer")
	}
	restApiErr := restApiServer.Run()
	logrus.Infof("Running http server on port %d", appConfig.Port)
	defer restApiServer.Shutdown()

	// Shutdown
	quit := notifyShutdown()
	select {
	case err := <-restApiErr:
		return errors.Wrap(err, "failed while running restApiServer")
	case <-quit:
		logrus.Info("Gracefully shutdown")
		return nil
	}
}

func notifyShutdown() chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	return quit
}
