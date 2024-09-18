package config

import "time"

type Config struct {
	Omdb OmdbConfig
	App  AppConfig
	Db   MysqlConfig
}

type AppConfig struct {
	Name string
	Port int
}

type OmdbConfig struct {
	BaseUrl            string
	ApiKey             string
	Timeout            time.Duration // in seconds
	MaxIdleConnPerHost int
}

type MysqlConfig struct {
	Host            string
	Port            string
	User            string
	Pass            string
	Schema          string
	MaxIdleConn     int
	MaxOpenConn     int
	MaxConnLifetime int
}
