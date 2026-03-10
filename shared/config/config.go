package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv     string
	ServerHost string

	ApiGatewayPort    string
	TenantServicePort string
	UserServicePort   string
	HRServicePort     string

	TenantDBUrl string
	UserDBUrl   string
	HRDBUrl     string

	RedisUrl    string
	RabbitMQUrl string

	JWTSecret        string
	JWTAccessTTLMin  int64
	JWTRefreshTTLMin int64
}

func Load() *Config {

	viper.AutomaticEnv()

	config := &Config{
		AppEnv:     viper.GetString("APP_ENV"),
		ServerHost: viper.GetString("SERVER_HOST"),

		ApiGatewayPort:    viper.GetString("API_GATEWAY_PORT"),
		TenantServicePort: viper.GetString("TENANT_SERVICE_PORT"),
		UserServicePort:   viper.GetString("USER_SERVICE_PORT"),
		HRServicePort:     viper.GetString("HR_SERVICE_PORT"),

		TenantDBUrl: viper.GetString("TENANT_DB_URL"),
		UserDBUrl:   viper.GetString("USER_DB_URL"),
		HRDBUrl:     viper.GetString("HR_DB_URL"),

		RedisUrl:    viper.GetString("REDIS_URL"),
		RabbitMQUrl: viper.GetString("RABBITMQ_URL"),

		JWTSecret:        viper.GetString("JWT_SECRET"),
		JWTAccessTTLMin:  viper.GetInt64("JWT_ACCESS_TTL_MIN"),
		JWTRefreshTTLMin: viper.GetInt64("JWT_REFRESH_TTL_MIN"),
	}

	return config
}
