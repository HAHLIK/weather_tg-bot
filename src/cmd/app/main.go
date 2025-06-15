package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("can't load .env file", err)
	}
	tgBotToken := os.Getenv("TG_BOT_TOKEN")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	tgBot, err := bot.New(tgBotToken, opts...)
	if err != nil {
		panic(err)
	}

	tgBot.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	switch update.Message.Text {

	case "/start":
		{
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Hello!",
			})
			break
		}

	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Добавлю потом прогноз погоды",
		})
	}
}
