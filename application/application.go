package application

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/maracko/go-torrent/application/torrent"
	"go.uber.org/zap"
)

type App struct {
	T *torrent.Torrent
	F fyne.App
	L *zap.SugaredLogger
}

func New(t *torrent.Torrent, l *zap.SugaredLogger) (*App, error) {
	return &App{
		T: t,
		F: app.New(),
		L: l,
	}, nil
}
