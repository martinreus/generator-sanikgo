package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// MustGetEnvInt gets an int environment variable or panics
func MustGetEnvInt(varName string) int {
	envInt, err := GetInt(varName)
	if err != nil {
		panic(err)
	}
	return envInt
}

// MustGetEnvString returns string from environment variable or panics if not available
func MustGetEnvString(varName string) string {
	val, err := GetString(varName)
	if err != nil {
		panic(err)
	}
	return val
}

func GetIntOrDefault(varName string, def int) int {
	if val, err := GetInt(varName); err == nil {
		return val
	}
	return def
}

func GetStringOrDefault(varName string, def string) string {
	if envVal, ok := os.LookupEnv(varName); ok {
		return envVal
	}
	return def
}

func GetInt(varName string) (int, error) {
	val, err := GetString(varName)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

func GetString(varName string) (string, error) {
	if val, ok := os.LookupEnv(varName); ok {
		return val, nil
	}
	return "", errors.New(fmt.Sprintf("environment variable %s not found", varName))
}
