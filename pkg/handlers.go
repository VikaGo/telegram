package pkg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"
const commandAbout = "about"
const commandLinks = "links"
const commandHelp = "help"
const commandHoliday = "holiday"

var holidayMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🇺🇸 USA"),
		tgbotapi.NewKeyboardButton("🇬🇧 UK"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🇫🇷 France"),
		tgbotapi.NewKeyboardButton("🇩🇪 Germany"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🇮🇹 Italy"),
		tgbotapi.NewKeyboardButton("🇮🇳 India"),
	),
)

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
	case commandHoliday:
		return b.handleHolidayCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessageText(message *tgbotapi.Message) error {

	switch message.Text {
	case "USA":
		reply := tgbotapi.NewMessage(message.Chat.ID, "You pressed button 1")
		b.bot.Send(reply)
	case "UK":
		reply := tgbotapi.NewMessage(message.Chat.ID, "You pressed button 1")
		b.bot.Send(reply)
	default:
		reply := tgbotapi.NewMessage(message.Chat.ID, "Please use the keyboard to choose an option.")
		reply.ReplyMarkup = holidayMenu
		b.bot.Send(reply)
	}

	return nil
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

func (b *Bot) handleHolidayCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Please choose a country:")
	msg.ReplyMarkup = holidayMenu
	_, err := b.bot.Send(msg)
	return err
}

//country := ""
//
//switch message.Text {
//case "🇺🇸 USA":
//	country = "US"
//case "🇬🇧 UK":
//	country = "GB"
//case "🇫🇷 France":
//	country = "FR"
//case "🇩🇪 Germany":
//	country = "DE"
//case "🇮🇹 Italy":
//	country = "IT"
//case "🇮🇳 India":
//	country = "IN"
//}
//
//if country != "" {
//	//holidays, err := getHolidays(country)
//	if err != nil {
//		msg := tgbotapi.NewMessage(message.Chat.ID, "An error occurred while getting holidays.")
//		b.bot.Send(msg)
//	}
//}

//else if len(holidays) == 0 {
//	msg := tgbotapi.NewMessage(message.Chat.ID, "There are no holidays today in "+country+".")
//	b.bot.Send(msg)
//}

//else {
//msgText := "Holidays today in " + country + ":\n"
//for _, holiday := range holidays {
//msgText += "- " + holiday.Name + "\n"
//}

//msg := tgbotapi.NewMessage(message.Chat.ID, msgText)}

//return err
//}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "The command is invalid")
	_, err := b.bot.Send(msg)
	return err
}

//func (b *Bot) handleMessageText(message *tgbotapi.Message) error {
//
//	switch message.Text {
//
//	case "USA":
//		// Handle the first button
//		reply := tgbotapi.NewMessage(message.Chat.ID, "You pressed button 1")
//		b.bot.Send(reply)
//	case "UK":
//		// Handle the second button
//		reply := tgbotapi.NewMessage(message.Chat.ID, "You pressed button 2")
//		b.bot.Send(reply)
//	case "France":
//		// Handle the third button
//		reply := tgbotapi.NewMessage(message.Chat.ID, "You pressed button 3")
//		b.bot.Send(reply)
//	case "Germany":
//		// Handle the fourth button
//		reply := tgbotapi.NewMessage(message.Chat.ID, "You pressed button 4")
//		b.bot.Send(reply)
//	}
//	return nil
//}
