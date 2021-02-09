package config

import (
	"time"
	"os"
	"fmt"
	_ "strings"
	"path/filepath"
	"github.com/spf13/viper"

)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("SONIC-TRANS-POST")
}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()

	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// global defaults

	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")


	// read config.yaml

	//viper.SetConfigFile(cfgFile)
	temp,err:= os.UserConfigDir();if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(temp)
	v.AddConfigPath(temp)
	//config.SetConfigFile()
	temp,err=os.Executable(); if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	exPath := filepath.Dir(temp)
	fmt.Println(exPath)
	v.AddConfigPath(exPath)


	temp,err=os.Getwd(); if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(temp)
	v.AddConfigPath(temp)


	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Errorf("Could not read config   viper error: %v", err)
	}

	var to=v.GetString("to")
	fmt.Println(to)






	return v
}
