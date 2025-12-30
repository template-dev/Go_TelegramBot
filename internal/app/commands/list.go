package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) ListCmd(inputMsg *tgbotapi.Message) {
	outputMsgText := "Here all the products are:\n"

	products := c.productService.GetProducts()
	for _, product := range products {
		outputMsgText += product.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)
	c.bot.Send(msg)
}
