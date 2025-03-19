package weather_searcher_mock

import (
	"fetch_sdet_geoloc-util/lib/types"
)

type MockWeatherSearcher struct {
	ThrowError bool
	Error      error
	Results    []types.SearchResult
	index      int
}

func (app *MockWeatherSearcher) GetByLocationName(input string) (types.SearchResult, error) {
	if app.ThrowError {
		return types.SearchResult{}, app.Error
	}
	app.index++
	return app.Results[app.index-1], nil
}

func (app *MockWeatherSearcher) GetByPostCode(input string) (types.SearchResult, error) {
	if app.ThrowError {
		return types.SearchResult{}, app.Error
	}
	app.index++
	return app.Results[app.index-1], nil
}
