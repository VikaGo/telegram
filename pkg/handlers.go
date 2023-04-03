package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"
const commandAbout = "about"
const commandLinks = "links"
const commandHelp = "help"


func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
	case commandAbout:
		return b.handleAboutCommand(message)
	case commandLinks:
		return b.handleLinksCommand(message)

	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't understand you. If you need /help, just choose the option.")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "/about \n/links \n/help ")
	_, err := b.bot.Send(msg)
	return err
}
func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I can help you \n/about \n/links \n/help")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleAboutCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hi there! My name is Vika and I'm 23 years old. I learn Go programming language.")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleLinksCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, " Facebook https://www.facebook.com/viktoria.kulyna.9 \nLinkedIn https://www.linkedin.com/in/vika-halenda-a48665155 \nGitHub https://github.com/VikaGo  ")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "The command is invalid")
	_, err := b.bot.Send(msg)
	return err
}

