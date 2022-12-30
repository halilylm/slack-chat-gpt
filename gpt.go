package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"log"
)

const (
	maxTokens   = 511
	temperature = 0.5
	engine      = gpt3.TextDavinci003Engine
)

// getAnswer gets answer from gpt
func getAnswer(apiKey, question string) string {
	var answer string
	i := 0
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	if err := client.CompletionStreamWithEngine(ctx, engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(maxTokens),
		Temperature: gpt3.Float32Ptr(temperature),
	}, func(resp *gpt3.CompletionResponse) {
		if i > 1 {
			answer += fmt.Sprint(resp.Choices[0].Text)
		}
		i++
	}); err != nil {
		log.Fatalln(err)
	}
	return answer
}
