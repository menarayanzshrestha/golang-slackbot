package main

import (
	"context"
	"fmt"
	"log"

	"github.com/menarayanzshrestha/slack-bot/utils"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events:")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	utils.LoadEnv()

	getEnv := utils.GetEnvWithKey

	slackBotToken := getEnv("SLACK_BOT_TOKEN")
	slackAppToken := getEnv("SLACK_APP_TOKEN")

	bot := slacker.NewClient(slackBotToken, slackAppToken)
	go printCommandEvents(bot.CommandEvents())
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
