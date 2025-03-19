package open_weather_map_wrapper

import (
	"net/http"
)

func NewOpenWeatherMapSearcher(client *http.Client, baseURL, credential string) *OpenWeatherMapSearcher {
	return &OpenWeatherMapSearcher{
		client:     client,
		credential: credential,
		baseURL:    baseURL,
	}
}

type OpenWeatherMapSearcher struct {
	client     *http.Client
	credential string
	baseURL    string
}
