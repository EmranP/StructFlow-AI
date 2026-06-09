package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV     string
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
	DBUrl  string

	AIGeminiKey  string
	AIClaudeKey  string
	AIChatGPTKey string

	SmtpHost     string
	SmtpPort     int
	SmtpEmail    string
	SmtpPassword string

	JWTSecret string
	OriginUrl string
}

func Load() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		ENV:       os.Getenv("ENV"),
		AppPort:   os.Getenv("APP_PORT"),
		OriginUrl: os.Getenv("CLIENT_URL"),

		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBUrl:  os.Getenv("DB_URL"),

		AIGeminiKey:  os.Getenv("API_AI_GEMINI_KEY"),
		AIClaudeKey:  os.Getenv("API_AI_CLAUDE_KEY"),
		AIChatGPTKey: os.Getenv("API_AI_CHAT_GPT_KEY"),

		SmtpHost:     os.Getenv("SMTP_HOST"),
		SmtpPort:     smtpPort,
		SmtpEmail:    os.Getenv("SMTP_EMAIL"),
		SmtpPassword: os.Getenv("SMTP_PASSWORD"),

		JWTSecret: os.Getenv("JWT_SECRET"),
	}, nil
}
