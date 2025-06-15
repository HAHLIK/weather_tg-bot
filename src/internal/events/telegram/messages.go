package telegramEvents

const (
	sorryMsg string = `
	Sorry, I don't understand the query. 
	You may have entered the wrong city name
	`

	helpMsg string = `
	I can send weather by your city.
	Just send me name your city.
	`

	helloMsg string = "Hello! \n" + helpMsg

	weatherMsg string = `
	Country: %s
	State: %s
	Location: %s
	Weather: %s
	Temp: %v ℃
	Pressure: %v
	Humidity: %v
	WindSpeed: %v m/s
	`
)
