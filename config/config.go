package config

import (
	"fmt"

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
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	appConfig := AppConfig{}
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name : ", appConfig.Name)

	return &appConfig
}

func (c *AppConfig) GetPort() string {
	return fmt.Sprintf(":%v", c.Port)
}
