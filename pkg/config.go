package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

// Config defines the structure of the config object
// it defines information about server details and
// postgres db details including the credentials
type Config struct {
	Server struct {
		Host string `json:"host" mapstructure:"host"`
		Port string `json:"port" mapstructure:"port"`
	} `json:"pkg" mapstructure:"pkg"`
	Db struct {
		Host string `json:"host" mapstructure:"host"`
		Port string `json:"port" mapstructure:"port"`
		Username string `json:"username" mapstructure:"username"`
		Password string `json:"password" mapstructure:"password"`
		Database string `json:"database" mapstructure:"database"`
	} `json:"db" mapstructure:"db"`
}

var config *Config

// InitializeConfig makes use of viper library to initialize
// config from multiple sources such as json, yaml, toml and
// even environment variables, it returns a pointer to Config
func InitializeConfig() *Config {
	mutex := sync.Mutex{}
	if config != nil {
		return config
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	viper.AddConfigPath("/opt/server")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("error reading config file, %s", err))
	}

	// Set undefined variables
	viper.SetDefault("pkg.host", "")
	viper.SetDefault("pkg.port", "5000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.username", "postgres")

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("unable to decode config file : %v", err))
	}
	return config
}


