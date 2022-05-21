package weather

import (
	"encoding/json"
	"io/ioutil"
)

type WeatherData struct {
	Weather      []WeatherDescription `json:"weather"`
	FormatedTemp `json:"main"`
	City         string `json:"name"`
}

type FormatedTemp struct {
	Temp       float32 `json:"temp"`
	Feels_like float32 `json:"feels_like"`
	Temp_min   float32 `json:"temp_min"`
	Temp_max   float32 `json:"temp_max"`
	Pressure   float32 `json:"pressure"`
	Humidity   float32 `json:"humidity"`
}

type WeatherDescription struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type WeatherInfo struct {
	FormatedTemp string
	Description  string
	City         string
}

type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

func LoadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	json.Unmarshal(bytes, &c)
	return c, nil
}
