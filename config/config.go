package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env"`
	HttpServer `yaml:"http_server"`
	Postgres   `yaml:"postgres"`
}

type HttpServer struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	JwtTokenSecret string `yaml:"jwt_token_secret"`
}

type Postgres struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	User    string `yaml:"user"`
	Passwd  string `yaml:"passwd"`
	DbName  string `yaml:"db_name"`
	SslMode string `yaml:"ssl_mode"`
}

var (
	config *Config
	once   sync.Once
)

func MustLoad() *Config {
	var err error
	once.Do(func() {
		config = &Config{}

		err = cleanenv.ReadConfig("config.yml", config)
	})
	if err != nil {
		panic(err)
	}

	return config
}