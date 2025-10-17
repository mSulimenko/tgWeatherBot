package router

const (
	messageUnknownCommand = "Sorry, i don`t know this command"
	messageBadRequest     = "Please, type request as in helper"
	messageBadCoords      = "Your coordinates should be correct"
	messageInternalError  = "Sorry, internal error with request"
)

const helpString = `Доступные команды:

/weather 55.75 37.61
Числа - координаты места
первое число широта, второе - долгота
`

type Router struct {
	weatherProvider WeatherProvider
	commands        map[string]handler
}

type WeatherProvider interface {
	GetWeatherByCoordinates(lat, lon string) (WeatherData, error)
}

type WeatherData struct {
	Type      string
	Temp      int
	FeelsLike int
	WindSpeed int
	Name      string
}
