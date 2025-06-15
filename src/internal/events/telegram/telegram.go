package telegramEvents

import (
	"context"
	"log"

	openweather "github.com/HAHLIK/weather_tg-bot/internal/app/clients/openweatherapi"
	"github.com/HAHLIK/weather_tg-bot/internal/events"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Processor struct {
	ow *openweather.Openweather
}

func New(ow *openweather.Openweather) *Processor {
	return &Processor{
		ow: ow,
	}
}

func (p *Processor) Fetch(update *models.Update) (events.Event, error) {
	chatId := update.Message.Chat.ID
	text := update.Message.Text

	event := events.Event{
		ChatID: chatId,
		Text:   text,
	}

	log.Printf("%s -> chatID: %v, text: %s",
		update.Message.From.Username,
		chatId,
		text)

	return event, nil
}

func (p *Processor) Process(ctx context.Context, b *bot.Bot, event events.Event) error {
	return p.doCmd(ctx, b, event.Text, event.ChatID)
}
