package config

import "github.com/infinity/infinity-service/server/config"

func LoadConfig() *config.Config {
	v := NewViper()

	var conf config.Config
	if err := v.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return &conf
}
