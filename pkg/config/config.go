package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	BaseDomain string             `mapstructure:"baseDomain"`
	HTTPServer HTTPServerConfig   `mapstructure:"httpServer"`
	GateServer []GateServerConfig `mapstructure:"gateServer"`
	GameServer GameServerConfig   `mapstructure:"gameServer"`
	Database   DatabaseConfig     `mapstructure:"database"`
}

type HTTPServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type GateServerConfig struct {
	Name  string `mapstructure:"name"`
	Title string `mapstructure:"title"`
	Addr  string `mapstructure:"addr"`
}

type GameServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

var DefaultConfig = Config{
	BaseDomain: "example.com",
	HTTPServer: HTTPServerConfig{
		Addr: "0.0.0.0:8080",
	},
	GateServer: []GateServerConfig{{
		Name:  "os_beta01",
		Title: "127.0.0.1:22102",
		Addr:  "127.0.0.1:22102",
	}},
	GameServer: GameServerConfig{
		Addr: "0.0.0.0:22102",
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
