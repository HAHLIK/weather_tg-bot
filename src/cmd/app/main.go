package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	openweather "github.com/HAHLIK/weather_tg-bot/internal/app/clients/openweatherapi"
	eventConsumer "github.com/HAHLIK/weather_tg-bot/internal/consumer/event-consumer"
	telegramEvents "github.com/HAHLIK/weather_tg-bot/internal/events/telegram"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("can't load .env file", err)
	}
	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	owApiKey := os.Getenv("OPENWEATHER_API_KEY")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	processor := telegramEvents.New(openweather.New(owApiKey))
	eventConsumer := eventConsumer.New(processor, processor)

	log.Print("Telegram bot is starting")
	err = eventConsumer.Start(ctx, tgBotToken)
	if err != nil {
		log.Fatal(err)
	}
}
