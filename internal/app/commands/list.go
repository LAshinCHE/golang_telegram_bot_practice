package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	products := c.productService.List()
	textAns := "Items:\n\n"
	for _, product := range products {
		textAns += product.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, textAns)

	c.bot.Send(msg)
}
