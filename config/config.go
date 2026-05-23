package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig
	Postgresql PostgresqlConfig
	Redis      RedisConfig
}

type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

type PostgresqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
	Sslmode  bool
	Driver   string
}

type RedisConfig struct {
	Addr         string
	Password     string
	Db           int
	MinIdleConns int
	PoolSize     int
	PoolTimeout  int
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName(filename)
	v.SetConfigType("yml")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	if err := v.Unmarshal(&c); err != nil {
		log.Printf("unable to decode config into struct: %v", err)
		return nil, err
	}

	return &c, nil
}
