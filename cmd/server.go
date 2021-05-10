package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/exuan/waka-api/api"
	"github.com/exuan/waka-api/api/router"
	"github.com/exuan/waka-api/internal/config"
	"github.com/exuan/waka-api/internal/database"
	"github.com/exuan/waka-api/internal/logger"
	"github.com/exuan/waka-api/model"
	"github.com/exuan/waka-api/repository"
	"github.com/exuan/waka-api/service"
	"github.com/urfave/cli/v2"
)

var (
	Server = &cli.Command{
		Name:        "server",
		Usage:       "server",
		Description: "waka-api server",
		Action:      runServer,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "addr",
				Aliases: []string{"a"},
				Value:   ":25000",
				Usage:   "listen address default :25000",
			},
			&cli.StringFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Value:   "develop",
				Usage:   "environment default develop",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config/config.toml",
				Usage:   "config path default ./config/config.toml",
			},
		},
	}
)

//@todo wire init
func runServer(c *cli.Context) (err error) {
	if api.Cfg, err = config.New(c.String("env"), c.Command.Name, c.String("addr"), c.String("config")); err != nil {
		return err
	}

	if api.Log, err = logger.New(api.Cfg, api.LogCfgs...); err != nil {
		return err
	}
	defer api.Log.Sync()

	// database
	dbCfs := make([]*database.Cfg, 0)
	if err := api.Cfg.UnmarshalKey("db", &dbCfs); err != nil {
		return err
	}

	if err := database.Connect(dbCfs); err != nil {
		return err
	}

	if err := database.DB(database.Default()).
		AutoMigrate(model.Migrations...); err != nil {
		return err
	}
	defer database.Close()

	api.Service = service.New(repository.New(database.DB(database.Default())))

	e := router.New()
	go func() {
		if err := e.Start(api.Cfg.SrvAddr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}
