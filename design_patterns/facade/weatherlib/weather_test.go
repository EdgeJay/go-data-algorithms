package weatherlib

import (
	"bytes"
	"io"
	"testing"
)

func getMockData() io.Reader {
	response := `{"coord":{"lon":103.87,"lat":1.41},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],"base":"stations","main":{"temp":305.54,"feels_like":309.5,"temp_min":304.15,"temp_max":307.15,"pressure":1008,"humidity":59},"visibility":10000,"wind":{"speed":2.1},"clouds":{"all":75},"dt":1588060499,"sys":{"type":1,"id":9479,"country":"SG","sunrise":1588028221,"sunset":1588072003},"timezone":28800,"id":1880294,"name":"Seletar","cod":200}`
	return bytes.NewReader([]byte(response))
}

func TestCurrentWeatherData_responseParser(t *testing.T) {
	reader := getMockData()
	weatherData := CurrentWeatherData{APIKey: ""}
	data, err := weatherData.responseParser(reader)
	expectedID := 1880294
	expectedName := "Seletar"

	if err != nil {
		t.Fatalf("Expected error to be nil, got %s", err.Error())
	}

	t.Logf("ID: %d", data.ID)
	t.Logf("Name: %s", data.Name)
	t.Logf("Weather[0].Main: %s", data.Weather[0].Main)
	t.Logf("Weather[0].Description: %s", data.Weather[0].Description)

	if data.ID != expectedID {
		t.Errorf("Mismatched ID, expected %d, got %d", expectedID, data.ID)
	}

	if data.Name != expectedName {
		t.Errorf("Mismatched name, expected %s, got %s", expectedName, data.Name)
	}
}
