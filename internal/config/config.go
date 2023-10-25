package config

import "github.com/spf13/viper"

type Config struct {
	Port          int32  `mapstructure:"port"`
	ApiURL        string `mapstructure:"api_url"`
	MongoUri      string `mapstructure:"mongo_uri"`
	MongoDatabase string `mapstructure:"mongo_database"`
}

func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
