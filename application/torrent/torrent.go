package torrent

import (
	"github.com/cenkalti/rain/torrent"
	"github.com/maracko/go-torrent/application/config"
)

type Torrent struct {
	Client *torrent.Session
	Config *config.GoTorrentConfig
}

func New(cfg *config.GoTorrentConfig) (*Torrent, error) {
	client, err := torrent.NewSession(cfg.ConvertToRainCFG())
	if err != nil {
		return nil, err
	}
	return &Torrent{Client: client, Config: cfg}, nil
}
