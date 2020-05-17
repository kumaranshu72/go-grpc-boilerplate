package config

import (
	"os"
	"strconv"
)

//GetEnv : get environment variable
func GetEnv(key, def string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = def
	}
	return value
}

// GetEnvInt : gets environment variable value in integer
func GetEnvInt(key string, def int) (value int, err error) {
	if tmp := os.Getenv(key); tmp != "" {
		value, err = strconv.Atoi(tmp)
	} else {
		value = def
	}
	return
}