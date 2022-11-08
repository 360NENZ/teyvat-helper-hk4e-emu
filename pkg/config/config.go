package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	LoggingDir string             `mapstructure:"loggingDir"`
	BaseDomain string             `mapstructure:"baseDomain"`
	AutoSignUp bool               `mapstructure:"autoSignUp"`
	PassSignIn bool               `mapstructure:"passSignIn"`
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
	Addr     string `mapstructure:"addr"`
	CertFile string `mapstructure:"certFile"`
	KeyFile  string `mapstructure:"keyFile"`
}

type GateServerConfig struct {
	Name        string `mapstructure:"name"`
	Title       string `mapstructure:"title"`
	Addr        string `mapstructure:"addr"`
	DispatchUrl string `mapstructure:"dispatchUrl"`
}

type FilterConfig struct {
	Enable bool     `mapstructure:"enable"`
	Rules  []string `mapstructure:"rules"`
}

type GameServerConfig struct {
	Enable bool         `mapstructure:"enable"`
	Addr   string       `mapstructure:"addr"`
	Filter FilterConfig `mapstructure:"filter"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

var DefaultConfig = Config{
	LoggingDir: "log",
	BaseDomain: "example.com",
	AutoSignUp: true,
	PassSignIn: false,
	HTTPServer: HTTPServerConfig{
		Enable: true,
		Addr:   "0.0.0.0:8080",
		TLS:    TLSConfig{Enable: false},
	},
	GateServer: []GateServerConfig{{
		Name:        "os_beta01",
		Title:       "127.0.0.1:22101",
		DispatchUrl: "http://127.0.0.1:8080/query_cur_region",
	}, {
		Name:  "os_beta02",
		Title: "127.0.0.1:22102",
		Addr:  "127.0.0.1:22102",
	}},
	GameServer: GameServerConfig{
		Enable: true,
		Addr:   "0.0.0.0:22102",
		Filter: FilterConfig{
			Enable: true,
			Rules: []string{
				"allow:*",
				"block:WindSeedClientNotify",
				"block:PlayerLuaShellNotify",
			},
		},
	},
	Database: DatabaseConfig{
		Driver: "sqlite",
		DSN:    "file:data/hk4e-emu.db?cache=shared&mode=rwc",
	},
}

func LoadConfig() (cfg Config) { return LoadConfigName("config") }

func LoadConfigName(name string) (cfg Config) {
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/hk4e-emu")
	viper.AddConfigPath("$HOME/.hk4e-emu")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("Config file not found, using the default config")
			cfg = DefaultConfig
			initLogger(cfg.LoggingDir)
			return
		} else {
			log.Panic().Err(err).Msg("Failed to read config file")
		}
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panic().Err(err).Msg("Failed to decode config file")
	}
	if tls := &cfg.HTTPServer.TLS; tls.Enable {
		if tls.CertFile == "" {
			tls.CertFile = "data/tls_cert.pem"
		}
		if tls.KeyFile == "" {
			tls.KeyFile = "data/tls_key.pem"
		}
	}
	initLogger(cfg.LoggingDir)
	return
}

func initLogger(dir string) {
	if dir != "" {
		log.Logger = log.Output(io.MultiWriter(zerolog.ConsoleWriter{Out: os.Stderr}, newRollingFile(dir))).With().Caller().Logger()
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
	}
}

func newRollingFile(dir string) io.Writer {
	if err := os.MkdirAll(dir, 0744); err != nil {
		log.Error().Err(err).Str("path", "log").Msg("can't create log directory")
		return nil
	}
	return &lumberjack.Logger{
		Filename: path.Join(dir, fmt.Sprintf("hk4e-emu-%s.log", time.Now().Format("2006-01-02"))),
	}
}
