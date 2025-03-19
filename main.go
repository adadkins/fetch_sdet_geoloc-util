package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	open_weather "fetch_sdet_geoloc-util/lib/open_weather_map_wrapper"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	values := getParams()

	openMapsCredential := os.Getenv("OPEN_MAPS_API_KEY")
	openMapsBaseURL := os.Getenv("OPEM_MAPS_BASE_URL")

	client := &http.Client{}

	wrapper := open_weather.NewOpenWeatherMapSearcher(client, openMapsBaseURL, openMapsCredential)
	app := NewApp(wrapper)

	searchResults := app.SearchLocationNameOrPostCode(values)

	for _, v := range searchResults {
		fmt.Printf(
			"Input:    %s\n"+
				"Name:     %s\n"+
				"Lat:      %.6f\n"+
				"Lon:      %.6f\n"+
				"Country:  %s\n"+
				"State:    %s\n\n",
			v.SearchInput, v.Name, v.Lat, v.Lon, v.Country, v.State,
		)
	}

}

func getParams() []string {
	if len(os.Args) < 2 {
		fmt.Println("Usage: geoloc-util <values...>")
		fmt.Println(`Example: ./geoloc-util "90292" "90034" "Chicago, IL"`)
		os.Exit(1)
	}

	var params []string
	for _, arg := range os.Args[1:] {
		params = append(params, strings.TrimSpace(arg))
	}

	return params
}
