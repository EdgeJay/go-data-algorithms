package weatherlib

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// CurrentWeatherDataRetriever provides methods to retrieve weather data of location
type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (WeatherData, error)
	GetByGeoCoordinates(lat, lon float32) (WeatherData, error)
}

// CurrentWeatherData implements CurrentWeatherDataRetriever and provides methods to retrieve weather data of location
type CurrentWeatherData struct {
	APIKey string
}

// GetByCityAndCountryCode returns weather data by looking up using city and country code
func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (*WeatherData, error) {
	return c.doRequest(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s&units=metric", city, countryCode, c.APIKey))
}

// GetByGeoCoordinates returns weather data by looking up using lat and lon
func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (*WeatherData, error) {
	return c.doRequest(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric", lat, lon, c.APIKey))
}

func (c *CurrentWeatherData) doRequest(uri string) (weather *WeatherData, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		byt, errMsg := ioutil.ReadAll(res.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Request failed, status code was %d. Error message: %s", res.StatusCode, errMsg)
		return
	}

	weather, err = c.responseParser(res.Body)
	res.Body.Close()
	return
}

func (c *CurrentWeatherData) responseParser(reader io.Reader) (*WeatherData, error) {
	wData := WeatherData{}
	err := json.NewDecoder(reader).Decode(&wData)

	if err != nil {
		return nil, err
	}

	return &wData, nil
}

// WeatherData contains weather data of location
type WeatherData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	}
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		OneHour    float32 `json:"1h"`
		ThreeHours float32 `json:"3h"`
	} `json:"rain,omitempty"`
	DT  int `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int `json:"timezone"`
}
