package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
	"time"
)

type BotStruct struct {
	bot      *tgbotapi.BotAPI
	logger   *slog.Logger
	cmdViews map[string]ViewFunc
}

func NewTgBotInitialize(bot *tgbotapi.BotAPI, lg *slog.Logger) *BotStruct {
	return &BotStruct{bot: bot, logger: lg}
}

func (b *BotStruct) Start(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		b.logger.Error("GetUpdatesChan error", slog.String("error", err.Error()))
		return err
	}

	for {
		select {
		case update := <-updates:
			updateCtx, updateCancel := context.WithTimeout(ctx, 4*time.Second)
			b.handleUpdate(updateCtx, update)
			updateCancel()
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}
