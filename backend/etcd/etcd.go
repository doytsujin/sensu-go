// Package etcd manages the embedded etcd server that Sensu uses for storing
// state consistently across sensu-backend processes.
//
// To use the embedded Etcd, you must first call NewEtcd(). This will configure
// and start Etcd and ensure that it starts correctly. The channel returned by
// Err() should be monitored--these are terminal/fatal errors for Etcd.
package etcd

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/embed"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/pkg/capnslog"
	"google.golang.org/grpc/grpclog"
)

const (
	// StateDir is the base path for Sensu's local storage.
	StateDir = "/var/lib/sensu"
	// EtcdStartupTimeout is the amount of time we give the embedded Etcd Server
	// to start.
	EtcdStartupTimeout = 60 // seconds
	// ClientListenURL is the default listen address for clients.
	ClientListenURL = "http://127.0.0.1:2379"
	// PeerListenURL is the default listen address for peers.
	PeerListenURL = "http://127.0.0.1:2380"
	// InitialCluster is the default initial cluster
	InitialCluster = "default=http://127.0.0.1:2380"
	// DefaultNodeName is the default name for this cluster member
	DefaultNodeName = "default"
	// ClusterStateNew specifies this is a new etcd cluster
	ClusterStateNew = "new"
	// ClusterStateExisting specifies ths is an existing etcd cluster
	ClusterStateExisting = "existing"
)

func init() {
	clientv3.SetLogger(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))
}

// Config is a configuration for the embedded etcd
type Config struct {
	DataDir                 string
	Name                    string // Cluster Member Name
	ListenPeerURL           string
	ListenClientURLs        string
	InitialCluster          string
	InitialClusterState     string
	InitialClusterToken     string
	InitialAdvertisePeerURL string
	TLSConfig               *TLSConfig
}

// TLSConfig wraps Crypto TLSInfo
type TLSConfig struct {
	Info TLSInfo
	TLS  *tls.Config
}

// TLSInfo wraps etcd transport TLSInfo
type TLSInfo transport.TLSInfo

// NewConfig returns a pointer to an initialized Config object with defaults.
func NewConfig() *Config {
	c := &Config{}
	c.DataDir = StateDir
	c.ListenClientURLs = ClientListenURL
	c.ListenPeerURL = PeerListenURL
	c.InitialCluster = InitialCluster
	c.InitialClusterState = ClusterStateNew
	c.Name = DefaultNodeName
	c.InitialAdvertisePeerURL = PeerListenURL

	return c
}

func ensureDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if mkdirErr := os.MkdirAll(path, 0700); mkdirErr != nil {
				return mkdirErr
			}
		} else {
			return err
		}
	}
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return fmt.Errorf("path exists and is not directory - %s", path)
	}
	return nil
}

// Etcd is a wrapper around github.com/coreos/etcd/embed.Etcd
type Etcd struct {
	cfg          *Config
	etcd         *embed.Etcd
	loopbackURLs []string
}

// BackendID returns the ID of the etcd cluster member
func (e *Etcd) BackendID() (result string) {
	return e.etcd.Server.ID().String()
}

// NewEtcd returns a new, configured, and running Etcd. The running Etcd will
// panic on error. The calling goroutine should recover() from the panic and
// shutdown accordingly. Callers must also ensure that the running Etcd is
// cleanly shutdown before the process terminates.
//
// Callers should monitor the Err() channel for the running etcd--these are
// terminal errors.
func NewEtcd(config *Config) (*Etcd, error) {
	cfg := embed.NewConfig()

	cfg.Name = config.Name

	cfgDir := filepath.Join(config.DataDir, "etcd", "data")
	walDir := filepath.Join(config.DataDir, "etcd", "wal")
	cfg.Dir = cfgDir
	cfg.WalDir = walDir
	if err := ensureDir(cfgDir); err != nil {
		fmt.Println("error 1")
		return nil, err
	}
	if err := ensureDir(walDir); err != nil {
		fmt.Println("error 2")
		return nil, err
	}

	var loopbackAddrs []string
	var clientURLs []url.URL
	listenClientURLs := strings.Split(config.ListenClientURLs, ",")
	fmt.Printf("LISTEN CLIENT URLS: %v", listenClientURLs)
	for _, clientURL := range listenClientURLs {
		listenClientURL, err := url.Parse(clientURL)
		if err != nil {
			fmt.Println("error 3")
			return nil, err
		}
		clientURLs = append(clientURLs, *listenClientURL)

		// ensure we always listen on loopback, use https if we have
		// a tls configuration.
		if listenClientURL.Hostname() != "127.0.0.1" && listenClientURL.Hostname() != "localhost" {
			// ensure we always listen on loopback

			l, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				fmt.Println("error 4")
				return nil, err
			}
			if err = l.Close(); err != nil {
				fmt.Println("error 5")
				logger.Error(err)
			}

			addr, err := net.ResolveTCPAddr("tcp", l.Addr().String())
			if err != nil {
				fmt.Println("error 6")
				return nil, err
			}

			scheme := "http"
			if config.TLSConfig != nil {
				scheme = "https"
			}

			loopbackClientURL, _ := url.Parse(fmt.Sprintf("%s://127.0.0.1:%d", scheme, addr.Port))

			clientURLs = append(clientURLs, *loopbackClientURL)
			loopbackAddrs = append(loopbackAddrs, loopbackClientURL.String())
		} else {
			loopbackAddrs = append(loopbackAddrs, listenClientURL.String())
		}
	}
	fmt.Println(clientURLs)
	fmt.Println(loopbackAddrs)

	listenPeerURL, err := url.Parse(config.ListenPeerURL)
	if err != nil {
		fmt.Println("error 7")
		return nil, err
	}

	advertisePeerURL, err := url.Parse(config.InitialAdvertisePeerURL)
	if err != nil {
		fmt.Println("error 8")
		return nil, err
	}

	cfg.ACUrls = clientURLs
	cfg.APUrls = []url.URL{*advertisePeerURL}
	cfg.LCUrls = clientURLs
	cfg.LPUrls = []url.URL{*listenPeerURL}
	cfg.InitialClusterToken = config.InitialClusterToken
	cfg.InitialCluster = config.InitialCluster
	cfg.ClusterState = config.InitialClusterState

	// Every 5 minutes, we will prune all values in etcd to only their latest
	// revision.
	cfg.AutoCompactionMode = "revision"
	// This has to stay in ns until https://github.com/coreos/etcd/issues/9337
	// is resolved.
	cfg.AutoCompactionRetention = "1"
	// Default to 4G etcd size. TODO: make this configurable.
	cfg.QuotaBackendBytes = int64(4 * 1024 * 1024 * 1024)

	if config.TLSConfig != nil {
		cfg.ClientTLSInfo = (transport.TLSInfo)(config.TLSConfig.Info)
		cfg.PeerTLSInfo = (transport.TLSInfo)(config.TLSConfig.Info)
		cfg.ClientTLSInfo.ClientCertAuth = false
		cfg.PeerTLSInfo.ClientCertAuth = false
	}

	capnslog.SetFormatter(NewLogrusFormatter())

	fmt.Println(cfg)
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		fmt.Println("error 9")
		return nil, err
	}

	select {
	case <-e.Server.ReadyNotify():
		logger.Info("Etcd ready to serve client connections")
	case <-time.After(EtcdStartupTimeout * time.Second):
		e.Server.Stop()
		fmt.Println("error 10")
		return nil, fmt.Errorf("Etcd failed to start in %d seconds", EtcdStartupTimeout)
	}

	return &Etcd{config, e, loopbackAddrs}, nil
}

// Name returns the configured name for Etcd.
func (e *Etcd) Name() string {
	return e.cfg.Name
}

// Err returns the error channel for Etcd or nil if no etcd is started.
func (e *Etcd) Err() <-chan error {
	return e.etcd.Err()
}

// Shutdown will cleanly shutdown the running Etcd.
func (e *Etcd) Shutdown() error {
	etcdStopped := e.etcd.Server.StopNotify()
	e.etcd.Close()
	<-etcdStopped
	return nil
}

// NewClient returns a new etcd v3 client. Clients must be closed after use.
func (e *Etcd) NewClient() (*clientv3.Client, error) {
	var tlsCfg *tls.Config
	if e.cfg.TLSConfig != nil {
		tlsCfg = e.cfg.TLSConfig.TLS
	}

	listeners := e.etcd.Clients
	if len(listeners) == 0 {
		return nil, errors.New("no etcd client listeners found for server")
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   e.loopbackURLs,
		DialTimeout: 5 * time.Second,
		TLS:         tlsCfg,
	})
	if err != nil {
		return nil, err
	}

	return cli, nil
}

// Healthy returns Etcd status information.
func (e *Etcd) Healthy() bool {
	client, err := e.NewClient()
	if err != nil {
		return false
	}
	mapi := clientv3.NewMaintenance(client)
	_, err = mapi.Status(context.TODO(), e.cfg.ListenClientURLs)
	return err == nil
}

// LoopbackURLs returns the lookback URLs used by etcd
func (e *Etcd) LoopbackURLs() []string {
	return e.loopbackURLs
}
