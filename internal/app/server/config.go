package server

import "time"

type Config struct {
	HTTP     HTTPConfig     `yaml:"http"`
	Logging  LoggingConfig  `yaml:"logging"`
	Database DatabaseConfig `yaml:"database"`
}

type HTTPConfig struct {
	ListenHost   string        `yaml:"listen_host"   env:"HTTP_LISTEN_HOST"   env-default:"0.0.0.0"`
	ListenPort   string        `yaml:"listen_port"   env:"HTTP_LISTEN_PORT"   env-default:"2000"`
	ReadTimeout  time.Duration `yaml:"read_timeout"  env:"HTTP_READ_TIMEOUT"  env-default:"5s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"HTTP_WRITE_TIMEOUT" env-default:"5s"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"  env:"HTTP_IDLE_TIMEOUT"  env-default:"10s"`
}

type LoggingConfig struct {
	Env      string   `yaml:"env"      env:"LOGGING_ENV"      env-default:"prod"`
	Level    string   `yaml:"level"    env:"LOGGING_LEVEL"    env-default:"info"`
	Encoding string   `yaml:"encoding" env:"LOGGING_ENCODING" env-default:"json"`
	Paths    []string `yaml:"paths"    env:"LOGGING_PATHS"    env-default:"stdout"`
}

type DatabaseConfig struct {
	Host            string        `yaml:"host"              env:"DB_HOST"              env-default:"postgres"`
	Port            string        `yaml:"port"              env:"DB_PORT"              env-default:"5432"`
	DBName          string        `yaml:"db_name"           env:"DB_NAME"              env-default:"tododb"`
	User            string        `yaml:"user"              env:"DB_USER"              env-default:"postgres"`
	Password        string        `yaml:"password"          env:"DB_PASSWORD"          env-default:"postgres" json:"-"`
	MaxConns        int32         `yaml:"max_conns"         env:"DB_MAX_CONNS"         env-default:"25"`
	MinConns        int32         `yaml:"min_conns"         env:"DB_MIN_CONNS"         env-default:"5"`
	MaxConnIdle     time.Duration `yaml:"max_conn_idle"     env:"DB_MAX_CONN_IDLE"     env-default:"30m"`
	MaxConnLifetime time.Duration `yaml:"max_conn_lifetime" env:"DB_MAX_CONN_LIFETIME" env-default:"1h"`
}
