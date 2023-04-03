package configs

type Config struct {
	TelegramToken string `env:"TOKEN"`
	LogLevel      string `env:"LOG_LEVEL"`
	logFilePath   string `env:"LOG_FILE_PATH"`
	logFormat     string `env:"LOG_FORMAT"`
}
