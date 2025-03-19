package open_weather_map_wrapper

import (
	"encoding/json"
	"errors"
	"fetch_sdet_geoloc-util/lib/types"
	"fmt"
	"net/http"
	"net/url"
)

func (app *OpenWeatherMapSearcher) GetByPostCode(input string) (types.SearchResult, error) {
	route := "/geo/1.0/zip"

	if len(input) != 5 {
		return types.SearchResult{}, errors.New("invalid US zip code format")
	}

	params := url.Values{}
	params.Set("zip", input+",US")
	params.Set("appid", app.credential)

	fullURL := fmt.Sprintf("%s%s?%s", app.baseURL, route, params.Encode())

	resp, err := app.client.Get(fullURL)
	if err != nil {
		return types.SearchResult{}, fmt.Errorf("http get error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.SearchResult{}, errors.New("non-200 response received")
	}

	var result types.SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return types.SearchResult{}, fmt.Errorf("error decoding response: %w", err)
	}
	result.SearchInput = input
	return result, nil
}
