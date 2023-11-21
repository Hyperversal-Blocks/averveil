package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config interface {
	GetConfig() *config
}

func (c *config) GetConfig() *config {
	return c
}

type config struct {
	Chain struct {
		PrivateKey string `mapstructure:"PRIVATE_KEY" json:"chain"`
		Endpoint   string `mapstructure:"ENDPOINT" json:"endpoint"`
	} `mapstructure:"CHAIN" json:"chain"`
	Contracts struct {
		HBLOCK struct {
			Address string      `mapstructure:"ADDRESS" json:"address"`
			ABI     interface{} `mapstructure:"ABI" json:"abi"`
		} `mapstructure:"HBLOCK" json:"hblock"`
	} `mapstructure:"CONTRACTS" json:"contracts"`
	Logger struct {
		Level int    `mapstructure:"LEVEL" json:"level"`
		Env   string `mapstructure:"ENV" json:"env"`
	} `mapstructure:"LOGGER" json:"logger"`
	Swarm struct {
		Host string `mapstructure:"HOST" json:"host"`
		PORT string `mapstructure:"PORT" json:"port"`
	} `mapstructure:"SWARM" json:"swarm"`
	Server struct {
		Host    string `mapstructure:"HOST" json:"host"`
		PORT    string `mapstructure:"PORT" json:"port"`
		CORSAGE int    `mapstructure:"CORS_AGE" json:"cors"`
	} `mapstructure:"SERVER" json:"server"`
	JWT struct {
		Secret string `mapstructure:"SECRET" json:"secret"`
		Issuer string `mapstructure:"ISSUER" json:"issuer"`
		Expiry int64  `mapstructure:"EXPIRY" json:"expiry"`
	} `mapStructure:"JWT" json:"jwt"`
	Store struct {
		Path    string `mapstructure:"PATH" json:"path"`
		InMem   bool   `mapstructure:"IN_MEM" json:"inMem"`
		Logging bool   `mapstructure:"LOGGING" json:"logging"`
	} `mapstructure:"STORE" json:"store"`
	Arduino struct {
		Key string `mapstructure:"API_KEY" json:"key"`
	} `mapstructure:"ARDUINO" json:"arduino"`
	Google struct {
		Sheets string `mapstrucutre:"SHEETS" json:"sheets"`
	} `mapstructure:"GOOGLE" json:"google"`
}

func Init() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("env.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error discovering config: %w", err)
	}

	conf := config{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &conf, nil
}
