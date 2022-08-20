package config

import (
	"time"

	"github.com/cenkalti/rain/torrent"
)

var (
	defaultRainCFG = torrent.Config{
		// Session
		Database:                               "~/.go-torrent/session.db",
		DataDir:                                "~/.go-torrent/data",
		DataDirIncludesTorrentID:               false,
		Host:                                   "0.0.0.0",
		PortBegin:                              20000,
		PortEnd:                                30000,
		MaxOpenFiles:                           10240,
		PEXEnabled:                             true,
		ResumeWriteInterval:                    30 * time.Second,
		PrivatePeerIDPrefix:                    "-GoTorrent" + torrent.Version + "-",
		PrivateExtensionHandshakeClientVersion: "GoTorrent " + torrent.Version,
		BlocklistUpdateInterval:                24 * time.Hour,
		BlocklistUpdateTimeout:                 10 * time.Minute,
		BlocklistEnabledForTrackers:            true,
		BlocklistEnabledForOutgoingConnections: true,
		BlocklistEnabledForIncomingConnections: true,
		BlocklistMaxResponseSize:               100 << 20,
		TorrentAddHTTPTimeout:                  30 * time.Second,
		MaxMetadataSize:                        30 << 20,
		MaxTorrentSize:                         10 << 20,
		MaxPieces:                              64 << 10,
		DNSResolveTimeout:                      5 * time.Second,
		ResumeOnStartup:                        true,
		HealthCheckInterval:                    10 * time.Second,
		HealthCheckTimeout:                     60 * time.Second,
		FilePermissions:                        0o750,

		// RPC Server
		RPCEnabled:         false,
		RPCHost:            "127.0.0.1",
		RPCPort:            7246,
		RPCShutdownTimeout: 5 * time.Second,

		// Tracker
		TrackerNumWant:              200,
		TrackerStopTimeout:          5 * time.Second,
		TrackerMinAnnounceInterval:  time.Minute,
		TrackerHTTPTimeout:          10 * time.Second,
		TrackerHTTPPrivateUserAgent: "GoTorrent/" + torrent.Version,
		TrackerHTTPMaxResponseSize:  2 << 20,
		TrackerHTTPVerifyTLS:        true,

		// DHT node
		DHTEnabled:             true,
		DHTHost:                "0.0.0.0",
		DHTPort:                7246,
		DHTAnnounceInterval:    30 * time.Minute,
		DHTMinAnnounceInterval: time.Minute,
		DHTBootstrapNodes: []string{
			"router.bittorrent.com:6881",
			"dht.transmissionbt.com:6881",
			"router.utorrent.com:6881",
			"dht.libtorrent.org:25401",
			"dht.aelitis.com:6881",
		},

		// Peer
		UnchokedPeers:                3,
		OptimisticUnchokedPeers:      1,
		MaxRequestsIn:                250,
		MaxRequestsOut:               250,
		DefaultRequestsOut:           50,
		RequestTimeout:               20 * time.Second,
		EndgameMaxDuplicateDownloads: 20,
		MaxPeerDial:                  80,
		MaxPeerAccept:                20,
		ParallelMetadataDownloads:    2,
		PeerConnectTimeout:           5 * time.Second,
		PeerHandshakeTimeout:         10 * time.Second,
		PieceReadTimeout:             30 * time.Second,
		MaxPeerAddresses:             2000,
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

	defaultConfig = GoTorrentConfig{
		Database:                               defaultRainCFG.Database,
		DataDir:                                defaultRainCFG.DataDir,
		DataDirIncludesTorrentID:               defaultRainCFG.DataDirIncludesTorrentID,
		Host:                                   defaultRainCFG.Host,
		PortBegin:                              defaultRainCFG.PortBegin,
		PortEnd:                                defaultRainCFG.PortEnd,
		ResumeWriteInterval:                    defaultRainCFG.ResumeWriteInterval,
		BlocklistUpdateInterval:                defaultRainCFG.BlocklistUpdateInterval,
		BlocklistUpdateTimeout:                 defaultRainCFG.BlocklistUpdateTimeout,
		BlocklistEnabledForTrackers:            defaultRainCFG.BlocklistEnabledForTrackers,
		BlocklistEnabledForOutgoingConnections: defaultRainCFG.BlocklistEnabledForOutgoingConnections,
		BlocklistEnabledForIncomingConnections: defaultRainCFG.BlocklistEnabledForIncomingConnections,
		ResumeOnStartup:                        defaultRainCFG.ResumeOnStartup,
		HealthCheckInterval:                    defaultRainCFG.HealthCheckInterval,
		HealthCheckTimeout:                     defaultRainCFG.HealthCheckTimeout,
		TrackerStopTimeout:                     defaultRainCFG.TrackerStopTimeout,
		DHTEnabled:                             defaultRainCFG.DHTEnabled,
		DHTHost:                                defaultRainCFG.DHTHost,
		DHTPort:                                defaultRainCFG.DHTPort,
		DHTAnnounceInterval:                    defaultRainCFG.DHTAnnounceInterval,
		DHTMinAnnounceInterval:                 defaultRainCFG.DHTMinAnnounceInterval,
		UnchokedPeers:                          defaultRainCFG.UnchokedPeers,
		OptimisticUnchokedPeers:                defaultRainCFG.OptimisticUnchokedPeers,
		MaxRequestsIn:                          defaultRainCFG.MaxRequestsIn,
		MaxRequestsOut:                         defaultRainCFG.MaxRequestsOut,
		DefaultRequestsOut:                     defaultRainCFG.DefaultRequestsOut,
		RequestTimeout:                         defaultRainCFG.RequestTimeout,
		MaxPeerDial:                            defaultRainCFG.MaxPeerDial,
		MaxPeerAccept:                          defaultRainCFG.MaxPeerAccept,
		ParallelMetadataDownloads:              defaultRainCFG.ParallelMetadataDownloads,
		PeerConnectTimeout:                     defaultRainCFG.PeerConnectTimeout,
		PeerHandshakeTimeout:                   defaultRainCFG.PeerHandshakeTimeout,
		PieceReadTimeout:                       defaultRainCFG.PieceReadTimeout,
		MaxPeerAddresses:                       defaultRainCFG.MaxPeerAddresses,
	}
)
