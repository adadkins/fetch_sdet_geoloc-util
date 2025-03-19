package open_weather_map_wrapper

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

func TestGetByLocationName(t *testing.T) {
	godotenv.Load("../../.env")

	apiKey := os.Getenv("OPEN_MAPS_API_KEY")
	baseURL := os.Getenv("OPEM_MAPS_BASE_URL")
	assert.NotEmpty(t, apiKey)
	assert.NotEmpty(t, baseURL)

	r, err := recorder.New("fixtures/openweather_locationname")
	assert.NoError(t, err)

	defer r.Stop()

	client := r.GetDefaultClient()
	searcher := NewOpenWeatherMapSearcher(client, baseURL, apiKey)

	t.Run("No Comma Input", func(t *testing.T) {
		//given
		input := "Los Angeles"

		//when
		result, err := searcher.GetByLocationName(input)

		//then
		assert.NoError(t, err)
		assert.NotEmpty(t, result.Name)
	})

	t.Run("With Comma Input", func(t *testing.T) {
		//given
		input := "Los Angeles, CA"

		//when
		result, err := searcher.GetByLocationName(input)

		//then
		assert.NoError(t, err)
		assert.NotEmpty(t, result.Name)
	})

	t.Run("Handles a non 200 response code", func(t *testing.T) {
		//given
		searcher.credential = "fake"
		input := "FakeCity"

		//when
		_, err := searcher.GetByLocationName(input)

		//then
		assert.Error(t, err)
	})
}
