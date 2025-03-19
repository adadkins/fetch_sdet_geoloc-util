package open_weather_map_wrapper

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestGetByPostCode(t *testing.T) {
	godotenv.Load("../../.env")

	apiKey := os.Getenv("OPEN_MAPS_API_KEY")
	baseURL := os.Getenv("OPEM_MAPS_BASE_URL")
	assert.NotEmpty(t, apiKey)
	assert.NotEmpty(t, baseURL)

	r, err := recorder.New("fixtures/openweather_postcode")
	assert.NoError(t, err)

	defer r.Stop()

	recorderClient := r.GetDefaultClient()
	searcher := NewOpenWeatherMapSearcher(recorderClient, baseURL, apiKey)

	t.Run("Invalid Zip Code", func(t *testing.T) {
		//given
		input := "123"

		//when
		_, err := searcher.GetByPostCode(input)

		//then
		assert.Error(t, err)
	})

	t.Run("Valid Zip Code", func(t *testing.T) {
		//given
		input := "90210"

		//when
		result, err := searcher.GetByPostCode(input)

		//then
		assert.NoError(t, err)
		assert.NotEmpty(t, result.Name)
	})

	t.Run("Handles a non 200 response code", func(t *testing.T) {
		//given
		searcher.credential = "fake"
		input := "00000"

		//when
		_, err := searcher.GetByPostCode(input)

		//then
		assert.Error(t, err)
	})
}
