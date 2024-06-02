package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log/slog"
	"runtime/debug"
)

type ViewFunc func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error

func (b *BotStruct) handleUpdate(ctx context.Context, update tgbotapi.Update) {

	defer func() {
		if p := recover(); p != nil {
			b.logger.Error("panic recovered", slog.String("error", p.(string)), slog.String("stacktrace", string(debug.Stack())))
		}
	}()

	if update.Message != nil || !update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		b.bot.Send(msg)
		return
	}

	var view ViewFunc

	cmd := update.Message.Command()

	cmdView, ok := b.cmdViews[cmd]
	if !ok {
		return
	}
	view = cmdView
	if err := view(ctx, b.bot, update); err != nil {
		b.logger.Error("handleUpdate error", err)
		if _, err := b.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при ввыполнении комнады")); err != nil {
			b.logger.Error("handleUpdate bot.Send error", err)
		}
	}
}
