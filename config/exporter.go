package config

import (
	"errors"
	"fmt"
	"live-tracking/config/reader"
	"live-tracking/pkg/model"
	"log"
	"sync"

	"github.com/spf13/viper"
)

const (
	envLabel     = "ENVIRONMENT"
	envDevLabel  = "DEVELOPMENT"
	envProdLabel = "PRODUCTION"

	envGrpcPortLabel = "GRPC_PORT"
	envHttpPortLabel = "HTTP_PORT"
	envLogLevelLabel = "LOG_LEVEL"
	envLogTimeFormatLabel = "LOG_TIME_FORMAT"

	grpcPortLabel = "grpc-port"
	httpPortLabel = "http-port"
	logLevelLabel = "log-level"
	logTimeFormatLabel = "log-time-format"
)

var (
	//Config : instance for exporting application configuration
	Config *model.Configuration
	once   sync.Once
)

// GetConfig will return viper configuration
func GetConfig(configPath string) *model.Configuration {
	once.Do(func() {
		Config = setConfig(reader.GetConfig(configPath))
		err := allConfigMembersHaveValue(Config)
		if err != nil {
			log.Fatalf("config for field found")
		}
	})
	return Config
}

// allConfigMembersHaveValue checks the given config have all necessary config fields
func allConfigMembersHaveValue(config *model.Configuration) error {
	if &config.GRPCPort == nil || &config.HTTPPort == nil ||
		config.Environment == "" {
		return errors.New("necessary fields not found")
	}
	return nil
}

// keyExistsInConfig will check given configuration header isn't empty.
func keyExistsInConfig(key string, m map[string]interface{}) error {
	err := errors.New("key not found")
	if len(m) == 0 {
		return err
	}
	return nil
}

// set configuration to the configuration
func setConfig(cfg *viper.Viper) *model.Configuration {
	c := new(model.Configuration)
	var value interface{}
	var interr error
	env := GetEnv(envLabel, envDevLabel)

	if err := keyExistsInConfig(env, cfg.GetStringMap(env)); err != nil {
		log.Fatalf("config for " + env + " not found")
	}
	c.Environment = env

	value = cfg.GetString(env + "." + grpcPortLabel)
	if c.GRPCPort = GetEnv(envGrpcPortLabel, value.(string)); c.GRPCPort == "" {
		log.Fatalf("error on parsing configuration file." + envGrpcPortLabel + " config for field found")
	}
	value = cfg.GetString(env + "." + httpPortLabel)
	if c.HTTPPort = GetEnv(envHttpPortLabel, value.(string)); c.HTTPPort == "" {
		log.Fatalf("error on parsing configuration file." + envHttpPortLabel + " config for field found")
	}
	value = cfg.GetInt(env + "." + logLevelLabel)
	if c.LogLevel, interr = GetEnvInt(envLogLevelLabel, value.(int)); interr != nil {
		log.Fatalf("error on parsing configuration file." + envLogLevelLabel + " config for field found")
	}
	value = cfg.GetString(env + "." + logTimeFormatLabel)
	if c.LogTimeFormat = GetEnv(envLogTimeFormatLabel, value.(string)); c.LogTimeFormat == "" {
		log.Fatalf("error on parsing configuration file."  +envLogTimeFormatLabel + " config for field found")
	}
	return c
}
