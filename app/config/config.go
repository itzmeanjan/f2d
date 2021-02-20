package config

import (
	"github.com/spf13/viper"
)

// Read - Reading .env file content, during application start up
func Read(file string) error {
	viper.SetConfigFile(file)

	return viper.ReadInConfig()
}

// Get - Get config value by key
func Get(key string) string {
	return viper.GetString(key)
}

// GetDbUser - Connection to be made as User
func GetDbUser() string {
	return Get("DB_USER")
}

// GetDbPassword - Password to be used when connecting to database
func GetDbPassword() string {
	return Get("DB_PASSWORD")
}

// GetDbHost - Database running on host, where to connection attempt to be made
func GetDbHost() string {
	return Get("DB_HOST")
}

// GetDbPort - Database port to connect to
func GetDbPort() string {
	return Get("DB_PORT")
}

// GetDbName - Database to connect to
func GetDbName() string {
	return Get("DB_NAME")
}
