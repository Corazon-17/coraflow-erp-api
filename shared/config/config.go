package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv     string
	ServerHost string

	ApiGatewayPort  string
	UserServicePort string
	HRServicePort   string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string

	UserDB string
	HRDB   string

	RedisHost string
	RedisPort string

	RabbitMQHost string
	RabbitMQPort string

	JWTSecret string
}

func Load() *Config {

	viper.AutomaticEnv()

	config := &Config{
		AppEnv:     viper.GetString("APP_ENV"),
		ServerHost: viper.GetString("SERVER_HOST"),

		ApiGatewayPort:  viper.GetString("API_GATEWAY_PORT"),
		UserServicePort: viper.GetString("USER_SERVICE_PORT"),
		HRServicePort:   viper.GetString("HR_SERVICE_PORT"),

		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresPort:     viper.GetString("POSTGRES_PORT"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),

		UserDB: viper.GetString("USER_DB"),
		HRDB:   viper.GetString("HR_DB"),

		RedisHost: viper.GetString("REDIS_HOST"),
		RedisPort: viper.GetString("REDIS_PORT"),

		RabbitMQHost: viper.GetString("RABBITMQ_HOST"),
		RabbitMQPort: viper.GetString("RABBITMQ_PORT"),

		JWTSecret: viper.GetString("JWT_SECRET"),
	}

	return config
}
