package main

import (
	"os"

	"github.com/LashinCHE/golang_test_bot/internal/app/commands"
	"github.com/LashinCHE/golang_test_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 30,
	}

	productService := product.NewService()
	commander := commands.NewCommander(bot, productService)
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		commander.HendlerUpdate(&update)
	}
}
