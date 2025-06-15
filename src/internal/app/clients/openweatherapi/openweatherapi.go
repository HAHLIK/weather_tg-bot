package wetheropenapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/HAHLIK/weather_tg-bot/internal/pkg"
)

const GeocodingRequestURL string = "http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=5&appid=%s"
const WeatherRequestURL string = "https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%s"

var ErrZeroLenghtResponce = errors.New("zero lenght responce")

type Openweather struct {
	apiKey string
}

func New(apiKey string) *Openweather {
	return &Openweather{
		apiKey: apiKey,
	}
}

func (w *Openweather) Coordinates(city string) (Coordinates, error) {
	responce, err := http.Get(fmt.Sprintf(
		GeocodingRequestURL,
		city,
		w.apiKey))

	if err != nil {
		return Coordinates{}, pkg.ErrorWrap("can't get coordinates", err)
	}

	if responce.StatusCode != http.StatusOK {
		return Coordinates{}, pkg.ErrorWrap("fail get coorditanes request", responce.StatusCode)
	}

	var geocodingResponce []GeocodingResponce
	err = json.NewDecoder(responce.Body).Decode(&geocodingResponce)
	if err != nil {
		return Coordinates{}, pkg.ErrorWrap("can't unmarshal responce", err)
	}

	if len(geocodingResponce) == 0 {
		return Coordinates{}, ErrZeroLenghtResponce
	}

	return Coordinates{
		Lat: geocodingResponce[0].Lat,
		Lon: geocodingResponce[0].Lon,
	}, nil
}

func (w *Openweather) Weather(coord Coordinates) (Weather, error) {
	responce, err := http.Get(fmt.Sprintf(
		WeatherRequestURL,
		coord.Lat,
		coord.Lon,
		w.apiKey))

	if err != nil {
		return Weather{}, pkg.ErrorWrap("can't get coordinates", err)
	}

	if responce.StatusCode != http.StatusOK {
		return Weather{}, pkg.ErrorWrap("fail get coorditanes request", responce.StatusCode)
	}

	var weatherResponce WeatherResponce
	err = json.NewDecoder(responce.Body).Decode(&weatherResponce)
	if err != nil {
		return Weather{}, pkg.ErrorWrap("can't unmarshal responce", err)
	}

	return Weather{
		Main:          weatherResponce.Weather[0].Main,
		TempInCelsius: pkg.CelsiumFromKelvin(weatherResponce.Main.TempInKelvin),
		Pressure:      weatherResponce.Main.Pressure,
		Humidity:      weatherResponce.Main.Humidity,
		WindSpeed:     weatherResponce.Wind.Speed,
	}, nil
}
