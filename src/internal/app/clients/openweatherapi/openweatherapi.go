package openweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
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

func (w *Openweather) Weather(city string) (Weather, error) {
	coord, err := w.coordinates(city)

	if err == ErrZeroLenghtResponce {
		return Weather{}, nil
	}

	if err != nil {
		return Weather{}, err
	}

	responce, err := http.Get(fmt.Sprintf(
		WeatherRequestURL,
		coord.Lat,
		coord.Lon,
		w.apiKey))

	if err != nil {
		return Weather{}, pkg.ErrorWrap("can't get weather", err)
	}

	if responce.StatusCode != http.StatusOK {
		return Weather{}, pkg.ErrorWrap("fail get weather request", responce.StatusCode)
	}

	var weatherResponce WeatherResponce
	err = json.NewDecoder(responce.Body).Decode(&weatherResponce)
	if err != nil {
		return Weather{}, pkg.ErrorWrap("can't unmarshal responce", err)
	}

	return Weather{
		Location:      coord.Location,
		Main:          weatherResponce.Weather[0].Main,
		TempInCelsius: math.Round(pkg.CelsiumFromKelvin(weatherResponce.Main.TempInKelvin)*10) / 10,
		Pressure:      weatherResponce.Main.Pressure,
		Humidity:      weatherResponce.Main.Humidity,
		WindSpeed:     weatherResponce.Wind.Speed,
	}, nil
}

func (w *Openweather) coordinates(city string) (Coordinates, error) {
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
		Location: Location{
			Country: geocodingResponce[0].Country,
			State:   geocodingResponce[0].State,
			Name:    geocodingResponce[0].Name,
		},
		Lat: geocodingResponce[0].Lat,
		Lon: geocodingResponce[0].Lon,
	}, nil
}
