package lifecycle

import (
	"github.com/sirupsen/logrus"
	"proj/public/env"
)

var Config = &Cfg{}

type Cfg struct {
	Http  HttpConfig  `env:"HTTP"`
	MySQL MySQLConfig `env:"MYSQL"`
	Redis RedisConfig `env:"REDIS"`
}

type HttpConfig struct {
	Port         int `env:"PORT,default=8080"`
	GraceTimeout int `env:"GRACE_TIMEOUT,default=5"`
	ReadTimeout  int `env:"READ_TIMEOUT,default=0"`
	WriteTimeout int `env:"WRITE_TIMEOUT,default=0"`
}

type MySQLConfig struct {
	Address         string `env:"ADDRESS,default=localhost:3306"`
	User            string `env:"USER,default=root"`
	Password        string `env:"PASSWORD,default=root"`
	DB              string `env:"DB,default=dev"`
	PoolMaxIdle     int    `env:"POOL_MAX_IDLE,default=10"`
	PoolMaxOpen     int    `env:"POOL_MAX_OPEN,default=100"`
	PoolMaxLifeTime int    `env:"POOL_MAX_LIFE_TIME,default=3600"`
}

type RedisConfig struct {
	Address     string `env:"ADDRESS,default=localhost:6379"`
	Password    string `env:"PASSWORD"`
	DB          int    `env:"DB,default=0"`
	PoolSize    int    `env:"POOL_SIZE,default=3"`
	MaxRetries  int    `env:"MAX_RETRIES,default=1"`
	MinIdleSize int    `env:"MIN_IDLE_SIZE,default=1"`
}

func init() {
	if err := env.ParseEnv(Config); err != nil {
		logrus.WithError(err).Panic("init config failure")
	}
}
