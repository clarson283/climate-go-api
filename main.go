package main

import (
	"fmt"

	"github.com/clarson283/climate-go-api/api"
	"github.com/clarson283/climate-go-api/configs"
	"github.com/rs/zerolog/log"
)

func main() {
	conf := configs.GetConfig()
	fmt.Println(conf)

	err := api.GetClimate(conf)
	if err != nil {
		log.Err(err).Msg("error returned from GetClimate()")
	}
}
