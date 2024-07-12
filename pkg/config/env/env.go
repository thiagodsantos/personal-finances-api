package env

import "os"

type EnvConfig struct {
	MySQL MySQLEnvConfig
}

var env *EnvConfig

type Env interface {
	LoadConfig()
	GetConfig() *EnvConfig
}

func LoadConfig() {
	env = &EnvConfig{
		MySQL: MySQLEnvConfig{
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Host:     os.Getenv("MYSQL_HOST"),
			DBName:   os.Getenv("MYSQL_DBNAME"),
		},
	}
}

func GetConfig() *EnvConfig {
	if env == nil {
		LoadConfig()
	}
	return env
}
