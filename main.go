package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandsEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}

func main(){
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	slack_bot_token := os.Getenv("SLACK_BOT_TOKEN")
	slack_app_token := os.Getenv("SLACK_APP_TOKEN")
	bot := slacker.NewClient(slack_bot_token, slack_app_token)

	go printCommandsEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter){
			year:=req.Param("year")
			yob, err:= strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2021 - yob 
			r:= fmt.Sprintf("age is %d", age)
			res.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = bot.Listen(ctx)
	if err != nil{
		log.Fatal(err)
	}
}