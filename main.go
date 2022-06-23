package main

import (
	"context"
	"fmt"
	"net/http"
	"oees/application"
	"oees/infrastructure/persistance"
	"oees/infrastructure/server"
	"oees/infrastructure/utilities"
	"oees/interfaces"
	"oees/interfaces/middlewares"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/go-hclog"
)

func main() {
	logger := utilities.NewConsoleLogger()
	serverConfig := utilities.NewServerConfig()

	repoStore, repoError := persistance.NewRepoStore(serverConfig, logger)
	if repoError != nil {
		logger.Error(repoError.Error())
		os.Exit(1)
	}

	logger.Info("Migrating Models...")
	migrationError := repoStore.Migrate()
	if migrationError != nil {
		logger.Error(migrationError.Error())
		os.Exit(1)
	}
	logger.Info("Model Migrations Done...")

	appStore := application.NewAppStore(repoStore)
	interfaceStore := interfaces.NewInterfaceStore(appStore, logger)
	middlewareStore := middlewares.NewMiddlewareStore(logger, appStore)
	httpServer := server.NewHTTPServer(*serverConfig, appStore, interfaceStore, middlewareStore)
	httpServer.Serve()

	server := http.Server{
		Addr:         ":" + serverConfig.ServerPort,
		Handler:      httpServer.Router,
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
		IdleTimeout:  300 * time.Second,
	}

	go func() {
		logger.Info(fmt.Sprintf("Starting Server at: %s...", serverConfig.ServerAddress+":"+serverConfig.ServerPort))

		err := server.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	logger.Info(fmt.Sprintf("Server Shutting Down... %s ", sig))

	//gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if shutDownError := server.Shutdown(ctx); shutDownError != nil {
		logger.Error(shutDownError.Error())
	}
	logger.Info("Server Shut Down.")
}
