package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type ServerConfig struct {
	AppVersion        string        `mapstructure:"appversion"`
	Port              string        `mapstructure:"port"`
	PprofPort         string        `mapstructure:"pprofport"`
	Mode              string        `mapstructure:"mode"`
	JwtSecretKey      string        `mapstructure:"jwtsecretkey"`
	CookieName        string        `mapstructure:"cookiename"`
	ReadTimeout       time.Duration `mapstructure:"readtimeout"`
	WriteTimeout      time.Duration `mapstructure:"writetimeout"`
	SSL               bool          `mapstructure:"ssl"`
	CtxDefaultTimeout time.Duration `mapstructure:"ctxdefaulttimeout"`
	CSRF              bool          `mapstructure:"csrf"`
	Debug             bool          `mapstructure:"debug"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Db       string `mapstructure:"db"`
	Sslmode  bool   `mapstructure:"sslmode"`
	Driver   string `mapstructure:"driver"`
}

type RedisConfig struct {
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	Db           int    `mapstructure:"db"`
	MinIdleConns int    `mapstructure:"minidleconns"`
	PoolSize     int    `mapstructure:"poolsize"`
	PoolTimeout  int    `mapstructure:"pooltimeout"`
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName(filename)
	v.SetConfigType("yml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	v.BindEnv("postgres.user", "POSTGRES_USER")
	v.BindEnv("postgres.password", "POSTGRES_PASSWORD")
	v.BindEnv("postgres.db", "POSTGRES_DB")

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
