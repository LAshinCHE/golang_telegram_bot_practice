package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/LashinCHE/golang_test_bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Можно обявить мапу и в качестве ключей хранить тег команды, а в качестве значения функцию команды
// var registeredCommands = map[string]func(c *Commander, message *tgbotapi.Message){}

var registeredCommands = map[string]func(c *Commander, message *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) HendlerUpdate(update *tgbotapi.Update) {

	if update.CallbackQuery != nil {
		parseData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parseData)

		log.Println("Try generate update message from callback query")
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed: %+v", parseData),
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil {
		return
	}

	command, ok := registeredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}

	// switch update.Message.Command() {
	// case "help":
	// 	c.Help(update.Message)
	// case "list":
	// 	c.List(update.Message)
	// case "get":
	// 	c.Get(update.Message)
	// default:
	// 	c.Default(update.Message)
	// }
}
