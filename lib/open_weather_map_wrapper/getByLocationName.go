package open_weather_map_wrapper

import (
	"encoding/json"
	"errors"
	"fetch_sdet_geoloc-util/lib/types"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (app *OpenWeatherMapSearcher) GetByLocationName(input string) (types.SearchResult, error) {
	route := "/geo/1.0/direct"
	if !strings.Contains(input, ",") {
		input = input + ",US"
	} else {
		input = input + ",US"
	}

	params := url.Values{}
	params.Set("q", input)
	params.Set("limit", "1")
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

	// The API returns an array of locations.
	var results []types.SearchResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return types.SearchResult{}, fmt.Errorf("error decoding response: %w", err)
	}

	if len(results) == 0 {
		return types.SearchResult{}, errors.New("no results found")
	}
	results[0].SearchInput = input

	return results[0], nil
}
