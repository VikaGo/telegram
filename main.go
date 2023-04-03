package main

import (
	"github.com/VikaGo/telegram.git/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type Config struct {
	TelegramToken string `env:"TOKEN"`
}

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	logFilePath := os.Getenv("LOG_FILE_PATH")
	logFormat := os.Getenv("LOG_FORMAT")

	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("Error parsing log level: %v", err)
	}

	logrus.SetLevel(level)

	if logFilePath != "" {
		file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Error opening log file: %v", err)
		}
		defer file.Close()
		logrus.SetOutput(file)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	if logFormat == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	TOKEN := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := pkg.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
		return
	}
}
