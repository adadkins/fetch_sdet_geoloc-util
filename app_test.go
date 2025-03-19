package main

import (
	"errors"
	"fetch_sdet_geoloc-util/lib/types"
	weather_searcher_mock "fetch_sdet_geoloc-util/lib/weather_search_mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmazonTrucks(t *testing.T) {
	t.Run("Can search for a Post code and City name", func(t *testing.T) {
		//given
		expectedResults := []types.SearchResult{
			{
				SearchInput: "90292",
				Name:        "Marina Del Rey",
				Lat:         41.883718,
				Lon:         -87.632382,
			},
			{
				SearchInput: "Chicago, IL",
				Name:        "Chicago",
				Lat:         41.883718,
				Lon:         -87.632382,
			},
		}
		searcher := &weather_searcher_mock.MockWeatherSearcher{
			Results: expectedResults,
		}
		app := NewApp(searcher)

		input := []string{"90292", "Chicago, IL"}

		//when
		result := app.SearchLocationNameOrPostCode(input)

		//then
		assert.Equal(t, expectedResults, result)
	})
	t.Run("Can handle when Searcher throws an error in PostCode search", func(t *testing.T) {
		//given
		expectedResults := []types.SearchResult{
			{
				SearchInput: "90292",
				Error:       "Error in PostCode Search",
			},
		}

		searcher := &weather_searcher_mock.MockWeatherSearcher{
			ThrowError: true,
			Error:      errors.New("Error in PostCode Search"),
		}
		app := NewApp(searcher)

		input := []string{"90292"}

		//when
		result := app.SearchLocationNameOrPostCode(input)

		//then
		assert.Equal(t, expectedResults, result)
	})
	t.Run("Can handle when Searcher throws an error in Location Name search", func(t *testing.T) {
		//given
		expectedResults := []types.SearchResult{
			{
				SearchInput: "Chicago, IL",
				Error:       "Error in LocationName Search",
			},
		}
		searcher := &weather_searcher_mock.MockWeatherSearcher{
			ThrowError: true,
			Error:      errors.New("Error in LocationName Search"),
		}
		app := NewApp(searcher)

		input := []string{"Chicago, IL"}

		//when
		result := app.SearchLocationNameOrPostCode(input)

		//then
		assert.Equal(t, expectedResults, result)
	})
}
