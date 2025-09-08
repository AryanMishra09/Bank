package util

import "github.com/spf13/viper"

// Config stores all configuration of the application
// The values are read by viper from a config file or environement variable
type Config struct {
	DBDriver      string `mapstructure:"dbDriver"`
	DBSource      string `mapstructure:"dbSource"`
	ServerAddress string `mapstructure:"serverAddress"`
}

// LoadConfig reads configurations from the file or environment files:
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
