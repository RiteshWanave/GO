package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents (analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main () {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5607464929509-5607568555493-8pBtssIEAwka7MHr9j1Y8bd0")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05HYDJ7ZC2-5610465011170-5515c61973c970e72ca9b09a94638f8cd55ec08aacd7076604072b0e36cc20c0")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents (bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition {
		Description: "yob calculator",
		// Examples: "my yob is 2023",
		Handler: func (botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023-yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}