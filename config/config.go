package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version int    `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
	Mock    []Mock `mapstructure:"mock"`
}

type Mock struct {
	Group   string             `mapstructure:"group"`
	Request []MockGroupRequest `mapstructure:"request"`
}

type MockGroupRequest struct {
	Path         string         `mapstructure:"path"`
	Method       string         `mapstructure:"method"`
	Status       int            `mapstructure:"status"`
	Type         string         `mapstructure:"type"`
	Body         string         `mapstructure:"body"`
	OverrideBody []OverrideBody `mapstructure:"override_body,omitempty"`
}

type OverrideBody struct {
	Condition OverrideCondition `mapstructure:"condition"`
	Status    int               `mapstructure:"status"`
	Type      string            `mapstructure:"type"`
	Body      string            `mapstructure:"body"`
}

type OverrideCondition struct {
	PayloadFrom  string      `mapstructure:"payload_from"`
	PayloadKey   string      `mapstructure:"payload_key"`
	PayloadValue interface{} `mapstructure:"payload_value"`
}

func Load() *AppConfig {
	viper.SetConfigType("yaml")

	// Load default config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.ReadInConfig()

	appConfig := AppConfig{}
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir("./mocks")
	if err != nil {
		panic(err)
	}

	// Load mocking rules
	viper.AddConfigPath("./mocks")
	for _, file := range files {
		fileName := file.Name()
		fileName = fileName[:len(fileName)-len(filepath.Ext(fileName))]

		viper.SetConfigName(fileName)
		viper.ReadInConfig()

		mock := Mock{}
		err := viper.Unmarshal(&mock)
		if err != nil {
			panic(err)
		}

		appConfig.Mock = append(appConfig.Mock, mock)
	}

	fmt.Println(fmt.Sprintf("\n%s mocks file loaded!", appConfig.Name))

	return &appConfig
}

func (c *AppConfig) GetPort() string {
	return fmt.Sprintf(":%v", c.Port)
}
