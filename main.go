package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"log"
)

func main() {
	// load env variables from .env file s
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// parse env variables to struct
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	// slack client
	client := slack.New(cfg.SlackAuthToken, slack.OptionAppLevelToken(cfg.SlackAppToken))

	// slack socket mode
	socket := socketmode.New(client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSlackEvent(ctx, &cfg, client, socket)

	socket.Run()
}
