package wetheropenapi

type GeocodingResponce struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type Coordinates struct {
	Lat float64
	Lon float64
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
	TempInKelvin float32 `json:"temp"`
	Pressure     float32 `json:"pressure"`
	Humidity     float32 `json:"humidity"`
}

type WindWeatherResponce struct {
	Speed float32 `json:"speed"`
}

type Weather struct {
	Main          string
	TempInCelsius float32
	Pressure      float32
	Humidity      float32
	WindSpeed     float32
}
