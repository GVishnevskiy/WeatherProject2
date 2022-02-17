package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func GetWeather(city string) (Weather, error) {
	rawUrl, exists := os.LookupEnv("WEATHER_API_URL")
	if !exists {
		LogData("No WEATHER_API_URL line in env file")
	}
	rawUrl += "/weather"
	apiKey, exists := os.LookupEnv("API_KEY")
	if !exists {
		LogData("No API_KEY line in env file, create one")
	}

	urlValues := url.Values{}
	urlValues.Add("q", city)
	urlValues.Add("appid", apiKey)
	urlValues.Add("units", "metric")

	url, err := url.Parse(rawUrl + "?" + urlValues.Encode())
	if LogErr(err) {
		return Weather{}, err
	}

	resp, err := http.Get(url.String())
	if LogErr(err) {
		return Weather{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	weather := Weather{}
	err = json.Unmarshal(body, &weather)
	if LogErr(err) {
		return Weather{}, err
	}

	return weather, nil
}
