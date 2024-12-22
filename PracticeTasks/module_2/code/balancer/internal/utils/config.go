package utils

import (
	"log"

	"github.com/spf13/viper"
)

// Config структура для хранения конфигурации приложения.
type Config struct {
	ServerPort string `mapstructure:"server_port"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	DBName     string `mapstructure:"db_name"`
}

// LoadConfig загружает конфигурацию из config.yaml или переменных окружения.
func LoadConfig(path string) (Config, error) {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv() // Для работы с переменными окружения

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Unable to decode config: %v", err)
		return config, err
	}
	return config, nil
}
