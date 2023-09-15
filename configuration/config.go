package configuration

import (
	"github.com/spf13/viper"
)

type Config struct {
	Chain struct {
		PrivateKey string `mapstructure:"PRIVATE_KEY"`
		Endpoint   string `mapstructure:"ENDPOINT"`
	} `mapstructure:"CHAIN"`
	Contracts struct {
		ExampleContract struct {
			Address string      `mapstructure:"ADDRESS"`
			ABI     interface{} `mapstructure:"ABI"`
		} `mapstructure:"EXAMPLE_CONTRACT"`
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
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

func Init() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("env.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = viper.Unmarshal(&config)

	return &config, nil
}
