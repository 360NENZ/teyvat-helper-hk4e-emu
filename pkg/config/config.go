package config

type Config struct {
	Database   DatabaseConfig     `mapstructure:"database"`
	GameServer GameServerConfig   `mapstructure:"game_server"`
	GateServer []GateServerConfig `mapstructure:"gate_server"`
	HTTPServer HTTPServerConfig   `mapstructure:"http_server"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type GameServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type GateServerConfig struct {
	Name  string `mapstructure:"name"`
	Title string `mapstructure:"title"`
	Addr  string `mapstructure:"addr"`
}

type HTTPServerConfig struct {
	Addr string `mapstructure:"addr"`
}

var DefaultConfig = Config{
	Database: DatabaseConfig{
		DSN: "file:data/sqlite3.db?cache=shared&mode=rwc",
	},
	GameServer: GameServerConfig{
		Addr: ":22102",
	},
	GateServer: []GateServerConfig{{
		Name:  "os_beta01",
		Title: "127.0.0.1:22102",
		Addr:  "127.0.0.1:22102",
	}},
	HTTPServer: HTTPServerConfig{
		Addr: ":8080",
	},
}
