package openweather

type GeocodingResponce struct {
	Country string  `json:"country"`
	State   string  `json:"state"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

type Coordinates struct {
	Location Location
	Lat      float64
	Lon      float64
}

type WeatherResponce struct {
	Weather []WeatherWeatherResponce `json:"weather"`
	Main    MainWeatherResponce      `json:"main"`
	Wind    WindWeatherResponce      `json:"wind"`
}

type WeatherWeatherResponce struct {
	Main string `json:"main"`
}

type MainWeatherResponce struct {
	TempInKelvin float64 `json:"temp"`
	Pressure     int     `json:"pressure"`
	Humidity     int     `json:"humidity"`
}

type WindWeatherResponce struct {
	Speed float32 `json:"speed"`
}

type Weather struct {
	Location      Location
	Main          string
	TempInCelsius float64
	Pressure      int
	Humidity      int
	WindSpeed     float32
}

type Location struct {
	Country string
	State   string
	Name    string
}
