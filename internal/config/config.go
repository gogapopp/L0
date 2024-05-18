package config

import (
	"flag"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MogrationDir string `yaml:"migration_dir"`
	HTTPServer   `yaml:"http_server"`
	Postgres     `yaml:"postgres"`
	Cache        `yaml:"cache"`
	Stan         `yaml:"stan"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Postgres struct {
	PostgresDb       string `yaml:"postgres_db"`
	PostgresUser     string `yaml:"postgres_user"`
	PostgresPassword string `yaml:"postgres_password"`
	PostgresHost     string `yaml:"postgres_host"`
	PostgresPort     string `yaml:"postgres_port"`
}

type Cache struct {
	CleanupInterval time.Time `yaml:"cleanup_interval"`
	DefaultTTL      time.Time `yaml:"default_ttl"`
}

type Stan struct {
	StanClusterID string `yaml:"cluster_id"`
	ClientID      string `yaml:"client_id"`
	DSN           string `yaml:"dsn"`
}

func New() (*Config, error) {
	configPath := flag.String("path", "config.yml", "path to config file")
	flag.Parse()

	var cfg Config

	if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
