package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	BaseDomain string             `mapstructure:"baseDomain"`
	AutoSignUp bool               `mapstructure:"autoSignUp"`
	HTTPServer HTTPServerConfig   `mapstructure:"httpServer"`
	GateServer []GateServerConfig `mapstructure:"gateServer"`
	GameServer GameServerConfig   `mapstructure:"gameServer"`
	Database   DatabaseConfig     `mapstructure:"database"`
}

type HTTPServerConfig struct {
	Enable bool      `mapstructure:"enable"`
	Addr   string    `mapstructure:"addr"`
	TLS    TLSConfig `mapstructure:"tls"`
}

type TLSConfig struct {
	Enable   bool   `mapstructure:"enable"`
	CertFile string `mapstructure:"certFile"`
	KeyFile  string `mapstructure:"keyFile"`
}

type GateServerConfig struct {
	Name        string `mapstructure:"name"`
	Title       string `mapstructure:"title"`
	Addr        string `mapstructure:"addr"`
	DispatchUrl string `mapstructure:"dispatchUrl"`
}

type GameServerConfig struct {
	Enable bool   `mapstructure:"enable"`
	Addr   string `mapstructure:"addr"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

var DefaultConfig = Config{
	BaseDomain: "example.com",
	HTTPServer: HTTPServerConfig{
		Enable: true,
		Addr:   "0.0.0.0:8080",
		TLS:    TLSConfig{Enable: false},
	},
	GateServer: []GateServerConfig{{
		Name:  "os_beta01",
		Title: "127.0.0.1:22102",
		Addr:  "127.0.0.1:22102",
	}},
	GameServer: GameServerConfig{
		Enable: true,
		Addr:   "0.0.0.0:22102",
	},
	Database: DatabaseConfig{
		Driver: "sqlite",
		DSN:    "file:data/hk4e-emu.db?cache=shared&mode=rwc",
	},
}

func init() {
	log.Logger = log.Output(io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, newRollingFile())).With().Caller().Logger()
}

func newRollingFile() io.Writer {
	if err := os.MkdirAll("log", 0744); err != nil {
		log.Error().Err(err).Str("path", "log").Msg("can't create log directory")
		return nil
	}
	return &lumberjack.Logger{
		Filename: path.Join("log", fmt.Sprintf("hk4e-emu-%s.log", time.Now().Format("2006-01-02-15-04-05"))),
	}
}

func LoadConfig() (cfg Config) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/hk4e-emu")
	viper.AddConfigPath("$HOME/.hk4e-emu")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("Config file not found, using the default config")
			cfg = DefaultConfig
			return
		} else {
			log.Panic().Err(err).Msg("Failed to read config file")
		}
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panic().Err(err).Msg("Failed to decode config file")
	}
	if tls := cfg.HTTPServer.TLS; tls.Enable {
		if tls.CertFile == "" {
			tls.CertFile = "data/tls_cert.pem"
		}
		if tls.KeyFile == "" {
			tls.KeyFile = "data/tls_key.pem"
		}
	}
	return
}
