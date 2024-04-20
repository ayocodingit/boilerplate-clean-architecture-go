package config

import (
	"log"

	"github.com/spf13/viper"
)

type Port struct {
	http int64
}

type App struct {
	Name string
	Env  string
	Port Port
	Log  string
}

type Pool struct {
	Min     int64
	Max     int64
	Acquire int64
	Idle    int64
}

type Db struct {
	Connection string
	Host       string
	Port       int64
	Username   string
	Password   string
	Name       string
	Pool       Pool
	KeepAlive  bool
}

type Jwt struct {
	AccessSecret string
	Algorithm    string
}

type Redis struct {
	Host string
	Port int64
	Ttl  int64
}

type Config struct {
	App   App
	Db    Db
	Jwt   Jwt
	Redis Redis
}

func NewConfig() Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error load .env file")
	}

	return Config{
		App: App{
			Name: viper.GetString("APP_NAME"),
			Env:  viper.GetString("APP_ENV"),
			Port: Port{
				http: viper.GetInt64("APP_PORT_HTTP"),
			},
			Log: viper.GetString("APP_LOG"),
		},
		Db: Db{
			Connection: viper.GetString("DB_CONNECTION"),
			Host:       viper.GetString("DB_HOST"),
			Port:       viper.GetInt64("DB_PORT"),
			Username:   viper.GetString("DB_USERNAME"),
			Password:   viper.GetString("DB_PASSWORD"),
			Name:       viper.GetString("DB_NAME"),
			Pool: Pool{
				Min:     viper.GetInt64("DB_POOL_MIN"),
				Max:     viper.GetInt64("DB_POOL_MAX"),
				Acquire: viper.GetInt64("DB_POOL_ACQUIRE"),
				Idle:    viper.GetInt64("DB_POOL_IDLE"),
			},
			KeepAlive: viper.GetBool("DB_KEEP_ALIVE"),
		},
		Jwt: Jwt{
			AccessSecret: viper.GetString("JWT_ACCESS_SECRET"),
			Algorithm:    viper.GetString("JWT_ALGORITHM"),
		},
		Redis: Redis{
			Host: viper.GetString("REDIS_HOST"),
			Port: viper.GetInt64("REDIS_PORT"),
			Ttl:  viper.GetInt64("REDIS_TTL"),
		},
	}
}
