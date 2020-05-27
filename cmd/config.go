package cmd

import (
	"github.com/spf13/viper"
)

const (
	defPort             = 8090
	defTillersSvcURL    = "tiller-deploy.kube-system.svc.cluster.local:44134"
	defEnableRemoteMode = false
	defPGEnabled        = true
	defPGURL            = "postgres.td-platform.svc.cluster.local"
	defPGUsername       = "postgres"
	defPGPassword       = "postgres"
	defPGPort           = 5432
	defRepoURL          = "localhost:8081"
	defRepoPath         = "/"
	defEnableLocalRepo  = false
	defClientCertPath   = "/etc/deploy-service.d/certs/client/tls.crt"
	defClientKeyPath    = "/etc/deploy-service.d/certs/client/tls.key"
	defClientCAPath     = "/etc/deploy-service.d/certs/client/ca.crt"
)

// Config is data structure to store service configuration
type Config struct {
	Port             int             `yaml:"port"`
	TillerServiceURL string          `yaml:"tillerServiceURL"`
	EnableRemoteMode bool            `yaml:"enableRemoteMode"`
	Postgres         PostgresConfig  `yaml:"postgres"`
	LocalRepo        LocalRepoConfig `yaml:"localRepo"`
	ClientTLS        TLSConfig       `yaml:"clientTLS"`
	viper            *viper.Viper
}

var config *Config

// PostgresConfig stores posgres access configuration
type PostgresConfig struct {
	Enabled  bool   `yaml:"enabled"`
	URL      string `yaml:"url"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// LocalRepoConfig stores local repo configuration
type LocalRepoConfig struct {
	Enabled bool   `yaml:"enabled"`
	Address string `yaml:"address"`
	Path    string `yaml:"path"`
}

// TLSConfig stores certificates for client or server
type TLSConfig struct {
	CA   string `yaml:"ca"`
	Cert string `yaml:"cert"`
	Key  string `yaml:"key"`
}

// NewConfig to get config with default options
func NewConfig() *Config {
	config := Config{
		Port:             defPort,
		TillerServiceURL: defTillersSvcURL,
		EnableRemoteMode: defEnableRemoteMode,
		Postgres: PostgresConfig{
			Enabled:  defPGEnabled,
			URL:      defPGURL,
			Port:     defPGPort,
			Username: defPGUsername,
			Password: defPGPassword,
		},
		LocalRepo: LocalRepoConfig{
			Enabled: defEnableLocalRepo,
			Address: defRepoURL,
			Path:    defRepoPath,
		},
		ClientTLS: TLSConfig{
			CA:   defClientCAPath,
			Cert: defClientCertPath,
			Key:  defClientKeyPath,
		},
	}
	return &config
}
