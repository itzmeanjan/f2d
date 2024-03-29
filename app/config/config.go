package config

import (
	"log"
	"strconv"

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
func GetDbPort() uint16 {

	port := Get("DB_PORT")

	_port, err := strconv.ParseUint(port, 10, 16)
	if err != nil {

		log.Printf("[❗️] Error when parsing database port number : %s\n", err.Error())
		return 0

	}

	return uint16(_port)

}

// GetDbName - Database to connect to
func GetDbName() string {
	return Get("DB_NAME")
}

// GetRPCHTTPURL - To be used for connecting to RPC node, over HTTP transport
func GetRPCHTTPURL() string {
	return Get("RPC_HTTP_URL")
}

// GetRPCWSURL - To be used for connecting to RPC node, over WS transport
func GetRPCWSURL() string {
	return Get("RPC_WS_URL")
}
