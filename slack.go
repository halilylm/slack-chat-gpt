package main

import (
	"context"
	"fmt"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"log"
)

// handleSlackEvent handles slack events
func handleSlackEvent(ctx context.Context, cfg *Config, client *slack.Client, socket *socketmode.Client) {
	for {
		select {
		case <-ctx.Done():
			log.Fatalln("shutting down socket mode listener")
		case event := <-socket.Events:
			switch event.Type {
			case socketmode.EventTypeEventsAPI:
				eventsAPI, ok := event.Data.(slackevents.EventsAPIEvent)
				if !ok {
					log.Printf("invalid event type")
					continue
				}
				socket.Ack(*event.Request)
				switch eventsAPI.Type {
				case slackevents.CallbackEvent:
					innerEvent := eventsAPI.InnerEvent
					switch ev := innerEvent.Data.(type) {
					case *slackevents.AppMentionEvent:
						if err := handleMention(cfg, ev, client); err != nil {
							log.Fatalln(err)
						}
					}
				default:
					log.Fatalln("unsupported event type")
				}
			}
		}
	}
}

// handleMention handles mentions to the bot and
// returns response from chat gpt
func handleMention(cfg *Config, event *slackevents.AppMentionEvent, client *slack.Client) error {
	question := event.Text
	answer := getAnswer(cfg.GptSecretKey, question)
	attachment := slack.Attachment{}
	attachment.Text = answer
	attachment.Color = cfg.BotColor
	if _, _, err := client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment)); err != nil {
		return fmt.Errorf("failed answer to the question due to: %w", err)
	}
	return nil
}
