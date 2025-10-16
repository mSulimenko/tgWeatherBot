package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	messageUnknownCommand = "Sorry, i don`t know this command"
)

type Router struct {
	commands map[string]handler
}

type WeatherProvider interface {
	GetWeatherByCoordinates(lat, lon float64) (WeatherData, error)
}

type WeatherData struct {
	// todo
}

type handler func(update tgbotapi.Update) string

func MakeRouter() *Router {
	commands := map[string]handler{
		"help":    commandHelpHandler,
		"weather": commandWeatherHandler,
	}

	return &Router{
		commands: commands,
	}
}

func (r *Router) Handle(update tgbotapi.Update) tgbotapi.MessageConfig {
	curHandler, ok := r.commands[update.Message.Command()]
	if !ok {
		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			messageUnknownCommand,
		)
	}

	respMessage := curHandler(update)

	return tgbotapi.NewMessage(
		update.Message.Chat.ID,
		respMessage,
	)
}

func commandHelpHandler(update tgbotapi.Update) string {
	return "helpHandler not done"
}

func commandWeatherHandler(update tgbotapi.Update) string {
	return "commandWeatherHandler not done"
}
