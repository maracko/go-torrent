package config

import (
	"io"
	"os"
	"time"

	"github.com/cenkalti/rain/torrent"
	intErr "github.com/maracko/go-torrent/application/errors"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var cfgPath string

func init() {
	cfgPath = os.Getenv("GOTORRENT_CONFIG")
	if cfgPath == "" {
		cfgPath = "./config.yaml"
	}
}

type GoTorrentConfig struct {
	// Session
	Database                               string        `json:"database,omitempty" yaml:"database"`
	DataDir                                string        `json:"dataDir,omitempty" yaml:"dataDir"`
	DataDirIncludesTorrentID               bool          `json:"dataDirIncludesTorrentID,omitempty" yaml:"dataDirIncludesTorrentID"`
	Host                                   string        `json:"host,omitempty" yaml:"host"`
	PortBegin                              uint16        `json:"portBegin,omitempty" yaml:"portBegin"`
	PortEnd                                uint16        `json:"portEnd,omitempty" yaml:"portEnd"`
	ResumeWriteInterval                    time.Duration `json:"resumeWriteInterval,omitempty" yaml:"resumeWriteInterval"`
	BlocklistUpdateInterval                time.Duration `json:"blocklistUpdateInterval,omitempty" yaml:"blocklistUpdateInterval"`
	BlocklistUpdateTimeout                 time.Duration `json:"blocklistUpdateTimeout,omitempty" yaml:"blocklistUpdateTimeout"`
	BlocklistEnabledForTrackers            bool          `json:"blocklistEnabledForTrackers,omitempty" yaml:"blocklistEnabledForTrackers"`
	BlocklistEnabledForOutgoingConnections bool          `json:"blocklistEnabledForOutgoingConnections,omitempty" yaml:"blocklistEnabledForOutgoingConnections"`
	BlocklistEnabledForIncomingConnections bool          `json:"blocklistEnabledForIncomingConnections,omitempty" yaml:"blocklistEnabledForIncomingConnections"`
	ResumeOnStartup                        bool          `json:"resumeOnStartup,omitempty" yaml:"resumeOnStartup"`
	HealthCheckInterval                    time.Duration `json:"healthCheckInterval,omitempty" yaml:"healthCheckInterval"`
	HealthCheckTimeout                     time.Duration `json:"healthCheckTimeout,omitempty" yaml:"healthCheckTimeout"`

	// Tracker
	TrackerStopTimeout time.Duration `json:"trackerStopTimeout,omitempty" yaml:"trackerStopTimeout"`

	// DHT node
	DHTEnabled             bool          `json:"dhtEnabled,omitempty" yaml:"dhtEnabled"`
	DHTHost                string        `json:"dhtHost,omitempty" yaml:"dhtHost"`
	DHTPort                uint16        `json:"dhtPort,omitempty" yaml:"dhtPort"`
	DHTAnnounceInterval    time.Duration `json:"dhtAnnounceInterval,omitempty" yaml:"dhtAnnounceInterval"`
	DHTMinAnnounceInterval time.Duration `json:"dhtMinAnnounceInterval,omitempty" yaml:"dhtMinAnnounceInterval"`

	// Peer
	UnchokedPeers             int           `json:"unchokedPeers,omitempty" yaml:"unchokedPeers"`
	OptimisticUnchokedPeers   int           `json:"optimisticUnchokedPeers,omitempty" yaml:"optimisticUnchokedPeers"`
	MaxRequestsIn             int           `json:"maxRequestsIn,omitempty" yaml:"maxRequestsIn"`
	MaxRequestsOut            int           `json:"maxRequestsOut,omitempty" yaml:"maxRequestsOut"`
	DefaultRequestsOut        int           `json:"defaultRequestsOut,omitempty" yaml:"defaultRequestsOut"`
	RequestTimeout            time.Duration `json:"requestTimeout,omitempty" yaml:"requestTimeout"`
	MaxPeerDial               int           `json:"maxPeerDial,omitempty" yaml:"maxPeerDial"`
	MaxPeerAccept             int           `json:"maxPeerAccept,omitempty" yaml:"maxPeerAccept"`
	ParallelMetadataDownloads int           `json:"parallelMetadataDownloads,omitempty" yaml:"parallelMetadataDownloads"`
	PeerConnectTimeout        time.Duration `json:"peerConnectTimeout,omitempty" yaml:"peerConnectTimeout"`
	PeerHandshakeTimeout      time.Duration `json:"peerHandshakeTimeout,omitempty" yaml:"peerHandshakeTimeout"`
	PieceReadTimeout          time.Duration `json:"pieceReadTimeout,omitempty" yaml:"pieceReadTimeout"`
	MaxPeerAddresses          int           `json:"maxPeerAddresses,omitempty" yaml:"maxPeerAddresses"`
}

// New initializes a config. If it doesn't exist one will be created with default settings.
// In case of error default config is always returned, and the error can be type switched for more information
func New() (*GoTorrentConfig, error) {
	out := &defaultConfig
	err := out.Load()
	return out, err
}

// Load attempts to read config file and if doesn't exist tries to write one with default settings
func (t *GoTorrentConfig) Load() error {
	f, err := os.Open(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			*t = defaultConfig
			if err = t.Save(); err != nil {
				return intErr.NewSaveError(err.Error())
			}
			return nil
		} else {
			return intErr.NewOpenError(err.Error())
		}
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "read error")
	}

	if err = yaml.Unmarshal(b, t); err != nil {
		return errors.Wrap(err, "unmarshal")
	}

	return nil
}

func (t *GoTorrentConfig) Save() error {
	var f *os.File
	f, err := os.Open(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = os.Create(cfgPath)
			if err != nil {
				intErr.NewSaveError(err.Error())
			}
		} else {
			return intErr.NewOpenError(err.Error())
		}
	}
	defer f.Close()

	b, err := yaml.Marshal(*t)
	if err != nil {
		return errors.Wrap(err, "marshal")
	}
	if _, err = f.Write(b); err != nil {
		return errors.Wrap(err, "write error")
	}

	return nil
}

func (t *GoTorrentConfig) ConvertToRainCFG() torrent.Config {
	if t == nil {
		return defaultRainCFG
	}

	return torrent.Config{
		// Session
		Database:                               t.Database,
		DataDir:                                t.DataDir,
		DataDirIncludesTorrentID:               t.DataDirIncludesTorrentID,
		Host:                                   t.Host,
		PortBegin:                              t.PortBegin,
		PortEnd:                                t.PortEnd,
		MaxOpenFiles:                           10240,
		PEXEnabled:                             true,
		ResumeWriteInterval:                    t.ResumeWriteInterval,
		PrivatePeerIDPrefix:                    "-GoTorrent" + torrent.Version + "-",
		PrivateExtensionHandshakeClientVersion: "GoTorrent " + torrent.Version,
		BlocklistUpdateInterval:                t.BlocklistUpdateInterval,
		BlocklistUpdateTimeout:                 t.BlocklistUpdateTimeout,
		BlocklistEnabledForTrackers:            t.BlocklistEnabledForTrackers,
		BlocklistEnabledForOutgoingConnections: t.BlocklistEnabledForOutgoingConnections,
		BlocklistEnabledForIncomingConnections: t.BlocklistEnabledForIncomingConnections,
		BlocklistMaxResponseSize:               100 << 20,
		TorrentAddHTTPTimeout:                  30 * time.Second,
		MaxMetadataSize:                        30 << 20,
		MaxTorrentSize:                         10 << 20,
		MaxPieces:                              64 << 10,
		DNSResolveTimeout:                      5 * time.Second,
		ResumeOnStartup:                        t.ResumeOnStartup,
		HealthCheckInterval:                    t.HealthCheckInterval,
		HealthCheckTimeout:                     t.HealthCheckTimeout,

		// RPC Server
		RPCEnabled:         false,
		RPCHost:            "127.0.0.1",
		RPCPort:            7246,
		RPCShutdownTimeout: 5 * time.Second,

		// Tracker
		TrackerNumWant:              200,
		TrackerStopTimeout:          t.TrackerStopTimeout,
		TrackerMinAnnounceInterval:  time.Minute,
		TrackerHTTPTimeout:          10 * time.Second,
		TrackerHTTPPrivateUserAgent: "GoTorrent/" + torrent.Version,
		TrackerHTTPMaxResponseSize:  2 << 20,
		TrackerHTTPVerifyTLS:        true,

		// DHT node
		DHTEnabled:             t.DHTEnabled,
		DHTHost:                t.DHTHost,
		DHTPort:                t.DHTPort,
		DHTAnnounceInterval:    t.DHTAnnounceInterval,
		DHTMinAnnounceInterval: t.DHTMinAnnounceInterval,
		DHTBootstrapNodes: []string{
			"router.bittorrent.com:6881",
			"dht.transmissionbt.com:6881",
			"router.utorrent.com:6881",
			"dht.libtorrent.org:25401",
			"dht.aelitis.com:6881",
		},

		// Peer
		UnchokedPeers:                t.UnchokedPeers,
		OptimisticUnchokedPeers:      t.OptimisticUnchokedPeers,
		MaxRequestsIn:                t.MaxRequestsIn,
		MaxRequestsOut:               t.MaxRequestsOut,
		DefaultRequestsOut:           t.DefaultRequestsOut,
		RequestTimeout:               t.RequestTimeout,
		EndgameMaxDuplicateDownloads: 20,
		MaxPeerDial:                  t.MaxPeerDial,
		MaxPeerAccept:                t.MaxPeerAccept,
		ParallelMetadataDownloads:    t.ParallelMetadataDownloads,
		PeerConnectTimeout:           t.PeerConnectTimeout,
		PeerHandshakeTimeout:         t.PeerHandshakeTimeout,
		PieceReadTimeout:             t.PieceReadTimeout,
		MaxPeerAddresses:             t.MaxPeerAddresses,
		AllowedFastSet:               10,

		// IO
		ReadCacheBlockSize: 128 << 10,
		ReadCacheSize:      256 << 20,
		ReadCacheTTL:       1 * time.Minute,
		ParallelReads:      1,
		ParallelWrites:     1,
		WriteCacheSize:     1 << 30,

		// Webseed settings
		WebseedDialTimeout:             10 * time.Second,
		WebseedTLSHandshakeTimeout:     10 * time.Second,
		WebseedResponseHeaderTimeout:   10 * time.Second,
		WebseedResponseBodyReadTimeout: 10 * time.Second,
		WebseedRetryInterval:           time.Minute,
		WebseedVerifyTLS:               true,
		WebseedMaxSources:              10,
		WebseedMaxDownloads:            4,
	}
}
