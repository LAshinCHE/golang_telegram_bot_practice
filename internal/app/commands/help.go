package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list {entities}\n"+
			"/get {number} - get product by number",
	)

	c.bot.Send(msg)
}

// Инициализируем команду
// Инит выполняется при загрузке модуля
// func init() {
// 	registeredCommands["help"] = (*Commander).Help
// }
