package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Conf struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.AddConfigPath(path)
	viper.SetConfigName(fmt.Sprintf("app.env.%s", os.Getenv("ENVIRONMENT")))
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, err
}
