package main

import (
	"order/internal/config"
	delivery "order/internal/delivery/http"
	"order/internal/repository"

	"fmt"
	"order/internal/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	lg "github.com/dmitryavdonin/gtools/logger"

	"github.com/dmitryavdonin/gtools/migrations"

	"github.com/dmitryavdonin/gtools/psql"
)

func main() {
	cfg, err := config.InitConfig("")
	if err != nil {
		panic(fmt.Sprintf("error initializing config %s", err))
	}

	//setup logger
	logger, err := lg.New(cfg.Log.Level, "order")
	if err != nil {
		panic(fmt.Sprintf("error initializing logger %s", err))
	}

	//db migrations
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	migrate, err := migrations.NewMigrations(dsn, "file://migrations")
	if err != nil {
		logger.Fatalf("migrations error: %s", err.Error())
	}

	err = migrate.Up()
	if err != nil {
		logger.Fatalf("migrations error: %s", err.Error())
	}

	//db init
	pg, err := psql.New(cfg.DB.Host, cfg.DB.Port, cfg.DB.Name, cfg.DB.User, cfg.DB.Password, psql.MaxPoolSize(cfg.DB.PoolMax), psql.ConnTimeout(time.Duration(cfg.DB.Timeout)*time.Second))
	if err != nil {
		logger.Fatal(fmt.Errorf("postgres connection error: %w", err))
	}

	//repository
	repository, err := repository.NewRepository(pg)
	if err != nil {
		logger.Fatal("storage initialization error: %s", err.Error())
	}

	//service
	services, err := service.NewServices(repository, logger)
	if err != nil {
		logger.Fatal("services initialization error: %s", err.Error())
	}

	delivery, err := delivery.New(services, cfg.App.Port, logger, delivery.Options{})
	if err != nil {
		logger.Fatal("delivery initialization error: %s", err.Error())
	}

	err = delivery.Run()
	if err != nil {
		logger.Fatal("start delivery error: %s", err.Error())
	}

	//closes connections on app kill
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	if err := shutdown(pg, logger); err != nil {
		logger.Fatal(fmt.Errorf("failed shutdown with error: %w", err))
	}

}

func shutdown(psql *psql.Postgres, logger *lg.Logger) error {
	fmt.Println("Gracefull shut down in progress...")
	psql.Pool.Close()
	logger.Info("Gracefull shutdown done!")
	return nil
}
