package reader

import (
	"sync"
	"log"
	"github.com/spf13/viper"
)

const (
	configType = "toml"
	configName = "config" // By default viper looks for config.toml file

)

var (
	config *viper.Viper
	once   sync.Once
)

// GetConfig if for fetching viper configs
func GetConfig(configPath string) *viper.Viper {
	once.Do(func() {
		config = viper.New()
		config.SetConfigName(configName)
		config.SetConfigType(configType)
		config.AddConfigPath(configPath)
		err := config.ReadInConfig()
		if err != nil {
			log.Fatalf("error on parsing configuration file %v", err)
		}
	})

	return config
}
