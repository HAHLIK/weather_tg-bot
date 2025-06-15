package telegramEvents

import (
	"context"
	"fmt"

	openweather "github.com/HAHLIK/weather_tg-bot/internal/app/clients/openweatherapi"
	"github.com/HAHLIK/weather_tg-bot/internal/pkg"
	"github.com/go-telegram/bot"
)

const (
	StartCmd = "/start"
	HelpCmd  = "/help"
)

func (p *Processor) doCmd(ctx context.Context, b *bot.Bot, text string, chatId int64) error {
	switch text {
	case StartCmd:
		return p.sendHello(ctx, b, chatId)
	case HelpCmd:
		return p.sendHelp(ctx, b, chatId)
	default:
		return p.sendWeather(ctx, b, chatId, text)
	}

}

func (p *Processor) sendHello(ctx context.Context, b *bot.Bot, chatId int64) error {
	err := p.sendMsg(ctx, b, chatId, helloMsg)
	if err != nil {
		return pkg.ErrorWrap("can't tg sendHello", err)
	}
	return nil
}

func (p *Processor) sendHelp(ctx context.Context, b *bot.Bot, chatId int64) error {
	err := p.sendMsg(ctx, b, chatId, helpMsg)
	if err != nil {
		return pkg.ErrorWrap("can't tg sendHelp", err)
	}
	return nil
}

func (p *Processor) sendWeather(ctx context.Context, b *bot.Bot, chatId int64, text string) (err error) {
	defer func() { err = pkg.ErrorWrap("can't tg sendWeather", err) }()

	if text == "" {
		p.sendMsg(ctx, b, chatId, sorryMsg)
		return err
	}

	weather, err := p.ow.Weather(text)
	if weather == (openweather.Weather{}) {
		p.sendMsg(ctx, b, chatId, sorryMsg)
		return err
	}

	message := fmt.Sprintf(weatherMsg,
		weather.Location.Country,
		weather.Location.State,
		weather.Location.Name,
		weather.Main,
		weather.TempInCelsius,
		weather.Pressure,
		weather.Humidity,
		weather.WindSpeed)
	p.sendMsg(ctx, b, chatId, message)

	return nil
}

func (p *Processor) sendMsg(ctx context.Context, b *bot.Bot, chatId int64, text string) error {
	_, err := b.SendMessage(
		ctx,
		&bot.SendMessageParams{
			ChatID: chatId,
			Text:   text,
		})
	return err
}
