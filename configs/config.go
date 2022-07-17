package configs

import (
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ClimateInterfaceUrl    string `hcl:"climate_interface_url"`
	ClimateInterfaceApiKey string `hcl:"climate_interface_api_key"`
}

func GetConfig() Config {
	var config Config
	err := hclsimple.DecodeFile("configs/config.hcl", nil, &config)
	if err != nil {
		log.Err(err).Msg("failed to load  config")
	}
	log.Info().Str("config url", config.ClimateInterfaceUrl).Msg("Config url")

	return config
}
