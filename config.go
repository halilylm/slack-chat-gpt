package main

type Config struct {
	SlackAuthToken string `env:"SLACK_AUTH_TOKEN,required"`
	SlackAppToken  string `env:"SLACK_APP_TOKEN,required"`
	GptSecretKey   string `env:"GPT_SECRET_KEY,required"`
	BotColor       string `env:"BOT_COLOR,required"`
}
