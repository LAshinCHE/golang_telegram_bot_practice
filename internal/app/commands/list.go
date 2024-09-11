package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	products := c.productService.List()
	textAns := "Items:\n\n"
	for _, product := range products {
		textAns += product.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, textAns)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
