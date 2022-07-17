package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/clarson283/climate-go-api/configs"
	"github.com/rs/zerolog/log"
)

func GetClimate(conf configs.Config) error {
	form := url.Values{}
	req, err := http.NewRequest("GET", conf.ClimateInterfaceUrl, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+conf.ClimateInterfaceApiKey)
	req.Header.Add("Content-Type", "application/application-json")

	resp, err := http.DefaultClient.Do(req)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("failed to read response body")
	}

	fmt.Println(string(body))

	return err
}
