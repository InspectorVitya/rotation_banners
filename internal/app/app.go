package app

import "github.com/inspectorvitya/rotation_banners/internal/interfaces"

type App struct {
	Logger  interfaces.Logger
	Storage interfaces.Storage
}

func New(logger interfaces.Logger, storage interfaces.Storage) *App {
	return &App{
		Logger:  logger,
		Storage: storage,
	}
}
