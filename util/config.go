package util

import "github.com/spf13/viper"

type Config struct {
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBSource       string `mapstructure:"DB_SOURCE"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	SECURITY_KEY   string `mapstructure:"SECURITY_KEY"`
	AccTokDuration string `mapstructure:"ACCESS_DURATION"`
	RefTokDuration string `mapstructure:"REFRESH_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
