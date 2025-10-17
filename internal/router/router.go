package router

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func (wd WeatherData) String() string {
	return fmt.Sprintf("В %s сейчас %s\nТемпература: %d°C (ощущается как %d°C)\nВетер: %d м/с",
		wd.Name, wd.Type, wd.Temp, wd.FeelsLike, wd.WindSpeed)
}

type handler func(update tgbotapi.Update) string

func MakeRouter(weatherProvider WeatherProvider) *Router {
	r := &Router{
		weatherProvider: weatherProvider,
	}
	commands := map[string]handler{
		"help":    r.commandHelpHandler,
		"weather": r.commandWeatherHandler,
	}

	r.commands = commands

	return r
}

func (r *Router) Handle(update tgbotapi.Update) tgbotapi.MessageConfig {
	curHandler, ok := r.commands[update.Message.Command()]
	if !ok {
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			messageUnknownCommand+"\n"+helpString,
		)
	}

	respMessage := curHandler(update)

	return tgbotapi.NewMessage(
		update.Message.Chat.ID,
		respMessage,
	)
}

func (r *Router) commandHelpHandler(update tgbotapi.Update) string {
	return helpString
}

func (r *Router) commandWeatherHandler(update tgbotapi.Update) string {
	req := strings.Fields(update.Message.Text)
	if len(req) != 3 {
		return messageBadRequest
	}
	lat, lon := req[1], req[2]

	if !validateLat(lat) || !validateLon(lon) {
		return messageBadCoords
	}

	weatherData, err := r.weatherProvider.GetWeatherByCoordinates(lat, lon)
	if err != nil {
		return messageInternalError
	}

	return weatherData.String()
}

func validateLat(lat string) bool {
	f, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return false
	}
	if f < -90 || f > 90 {
		return false
	}
	return true
}

func validateLon(lon string) bool {
	f, err := strconv.ParseFloat(lon, 64)
	if err != nil {
		return false
	}
	if f < -180 || f > 180 {
		return false
	}
	return true
}
