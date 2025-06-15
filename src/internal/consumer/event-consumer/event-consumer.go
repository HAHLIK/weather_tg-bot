package eventConsumer

import (
	"context"
	"log"

	"github.com/HAHLIK/weather_tg-bot/internal/events"
	"github.com/HAHLIK/weather_tg-bot/internal/pkg"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EventConsumer struct {
	Fetcher   events.Fetcher
	Processor events.Processor
}

func New(fetcher events.Fetcher, processor events.Processor) *EventConsumer {

	return &EventConsumer{
		Fetcher:   fetcher,
		Processor: processor,
	}
}

func (e *EventConsumer) Start(ctx context.Context, botToken string) error {
	opts := []bot.Option{
		bot.WithDefaultHandler(e.handler),
	}

	tgBot, err := bot.New(botToken, opts...)

	if err != nil {
		return pkg.ErrorWrap("can't start event consumer", err)
	}

	tgBot.Start(ctx)
	return nil
}

func (e *EventConsumer) handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	event, _ := e.Fetcher.Fetch(update)
	err := e.Processor.Process(ctx, b, event)

	if err != nil {
		log.Println(err)
	}
}
