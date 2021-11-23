package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/inspectorvitya/rotation_banners/internal/app"
	"github.com/inspectorvitya/rotation_banners/internal/configuration"
	"github.com/inspectorvitya/rotation_banners/internal/logger"
	httpserver "github.com/inspectorvitya/rotation_banners/internal/server/http"
	sqlstorage "github.com/inspectorvitya/rotation_banners/internal/storage"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "configuration-path", "configs/apiserver.toml", "path to configuration file")
}
func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := configuration.NewConfig(configPath)
	fmt.Println(cfg)
	if err != nil {
		log.Fatal(err)
	}

	loggerApp, err := logger.New(cfg.Logger)
	if err != nil {
		log.Fatal(err)
	}

	urlExample := "postgres://postgres:postgres@localhost:5432/rotbanner"
	db, err := sqlstorage.New(ctx, urlExample, loggerApp)

	if err != nil {
		log.Fatal(err)
	}

	App := app.New(loggerApp, db)
	s := httpserver.New(cfg, App)
	go signalHandler(ctx, s, db, cancel)
	go func() {
		if err := s.Start(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				App.Logger.Info("http server stopped")
			} else {
				App.Logger.Error("failed to start http server: " + err.Error())
				cancel()
			}
		}
	}()

	<-ctx.Done()
}

func signalHandler(ctx context.Context, server *httpserver.Server, db *sqlstorage.StorageInDB, cancel context.CancelFunc) {
	defer cancel()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	select {
	case <-signals:
		signal.Stop(signals)
		serverCloseCtx, serverCancel := context.WithTimeout(context.Background(), time.Second*3)
		defer serverCancel()

		if err := server.Close(serverCloseCtx); err != nil {
			server.App.Logger.Error("failed to stop http server: " + err.Error())
		}
		db.Close(serverCloseCtx)
	case <-ctx.Done():
	}
}
