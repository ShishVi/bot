package commands

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outMessage := "Here all products \n\n"

	products := c.productService.List()

	for _, item := range products {
		calories := strconv.Itoa(item.Calories)
		outMessage += item.Title + " " + calories
		outMessage += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMessage)
	c.bot.Send(msg)
}
