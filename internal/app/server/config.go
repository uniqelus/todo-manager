package server

import "time"

type Config struct {
	HTTP    HTTPConfig    `yaml:"http"`
	Logging LoggingConfig `yaml:"logging"`
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
