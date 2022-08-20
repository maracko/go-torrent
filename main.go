package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/maracko/go-torrent/application"
	"github.com/maracko/go-torrent/application/components/toolbar"
	"github.com/maracko/go-torrent/application/config"
	intErr "github.com/maracko/go-torrent/application/errors"
	"github.com/maracko/go-torrent/application/torrent"
	"go.uber.org/zap"
)

func main() {
	// Setup log
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	log := l.Sugar()
	//

	cfg, cfgErr := config.New()

	t, err := torrent.New(cfg)
	if err != nil {
		log.Fatalln("init torrent struct fail:", err)
	}

	app, err := application.New(t, log)
	if err != nil {
		log.Fatalln("init app struct fail:", err)
	}

	w := app.F.NewWindow("GO Torrent")
	w.SetContent(toolbar.Toolbar)
	if cfgErr != nil {
		handleCfgErr(app.F, cfgErr)
	}

	w.ShowAndRun()
}

func handleCfgErr(a fyne.App, cfgErr error) {
	var title, detail string
	switch cfgErr.(type) {
	case intErr.OpenError:
		title = "Open config error"
		detail = "An error occured while trying to open config file. Now using default settings.\nMore info: %v"
	case intErr.SaveError:
		title = "Save config error"
		detail = "An error occured while trying to save config file. Now using default settings.\nMore info: %v"
	default:
		title = "Config initialization error"
		detail = "Something went wrong while setting up configuration. Now using default settings.\nMore info: %v"
	}
	errW := a.NewWindow(title)
	errW.SetContent(widget.NewLabel(fmt.Sprintf(detail, cfgErr)))
	errW.Show()
}
