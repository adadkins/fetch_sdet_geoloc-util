package main

import (
	"fetch_sdet_geoloc-util/lib/types"
	"regexp"
)

type App struct {
	weatherSearcher types.WeatherSearcher
}

func NewApp(searcher types.WeatherSearcher) *App {
	return &App{
		weatherSearcher: searcher,
	}
}

func (app *App) SearchLocationNameOrPostCode(inputs []string) []types.SearchResult {
	results := []types.SearchResult{}
	for _, v := range inputs {
		if regexp.MustCompile(`^\d{5}$`).MatchString(v) {
			result, err := app.weatherSearcher.GetByPostCode(v)
			if err != nil {
				result.SearchInput = v
				result.Error = err.Error()
			}
			results = append(results, result)
			continue
		}
		result, err := app.weatherSearcher.GetByLocationName(v)
		if err != nil {
			result.SearchInput = v
			result.Error = err.Error()
		}
		results = append(results, result)
	}
	return results
}
