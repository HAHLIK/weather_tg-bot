package events

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Fetcher interface {
	Fetch(update *models.Update) (Event, error)
}

type Processor interface {
	Process(ctx context.Context, b *bot.Bot, event Event) error
}

type Event struct {
	Text   string
	ChatID int64
}
