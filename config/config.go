package config

import (
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env"`
	HttpServer  `yaml:"http_server"`
	Postgres    `yaml:"postgres"`
	Token       `yaml:"token"`
	ExternalApi `yaml:"external_api"`
}

type HttpServer struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
}

type Postgres struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	User    string `yaml:"user"`
	Passwd  string `yaml:"passwd"`
	DbName  string `yaml:"db_name"`
	SslMode string `yaml:"ssl_mode"`
}

type Token struct {
	MaxAge         int           `yaml:"maxage"`
	ExpiredIn      time.Duration `yaml:"expired_in"`
	JwtTokenSecret string        `yaml:"jwt_token_secret"`
}

type ExternalApi struct {
	NumverifyAccessKey string `yaml:"numverify_access_key"`
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
