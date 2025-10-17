package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tgWeatherBot/internal/config"
	rt "tgWeatherBot/internal/router"
	"tgWeatherBot/internal/weather"
)

const (
	configTempPath = "./configs/config.yaml"
)

func main() {

	cfg := config.MustReadConfig(configTempPath)

	bot, err := tgbotapi.NewBotAPI(cfg.TgAPIToken)
	if err != nil {
		panic(err)
	}

	weatherClient := weather.MakeClient(cfg.WeatherAPIKey)
	router := rt.MakeRouter(weatherClient)

	bot.Debug = false
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := router.Handle(update)
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

}
