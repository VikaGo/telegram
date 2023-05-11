package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}

}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		if update.Message.Text != "" {
			country := update.Message.Text
			holiday, err := getHoliday(country)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, an error occurred while retrieving the holiday information.")
				b.bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "In "+country+" is "+holiday+".")
				b.bot.Send(msg)
			}
			continue
		}
		b.handleMessage(update.Message)
	}

	return nil
