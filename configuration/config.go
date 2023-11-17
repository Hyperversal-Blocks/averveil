package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Chain struct {
		PrivateKey string `mapstructure:"PRIVATE_KEY"`
		Endpoint   string `mapstructure:"ENDPOINT"`
	} `mapstructure:"CHAIN"`
	Contracts struct {
		HBLOCK struct {
			Address string      `mapstructure:"ADDRESS"`
			ABI     interface{} `mapstructure:"ABI"`
		} `mapstructure:"HBLOCK"`
	} `mapstructure:"CONTRACTS"`
	Logger struct {
		Level int    `mapstructure:"LEVEL"`
		Env   string `mapstructure:"ENV"`
	} `mapstructure:"LOGGER"`
	Swarm struct {
		Host string `mapstructure:"HOST"`
		PORT string `mapstructure:"PORT"`
	} `mapstructure:"SWARM"`
	Server struct {
		Host    string `mapstructure:"HOST"`
		PORT    string `mapstructure:"PORT"`
		CORSAGE int    `mapstructure:"CORS_AGE"`
	} `mapstructure:"SERVER"`
	JWT struct {
		Secret string `mapstructure:"SECRET"`
		Issuer string `mapstructure:"ISSUER"`
		Expiry int64  `mapstructure:"EXPIRY"`
	} `mapStructure:"JWT"`
	Store struct {
		Path    string `mapstructure:"PATH"`
		InMem   bool   `mapstructure:"IN_MEM"`
		Logging bool   `mapstructure:"LOGGING"`
	} `mapstructure:"STORE"`
	Arduino struct {
		Key string `mapstructure:"API_KEY"`
	} `mapstructure:"ARDUINO"`
}

func Init() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("env.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error discovering config: %w", err)
	}

	config := Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}
