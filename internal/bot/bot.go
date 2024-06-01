package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"log/slog"
)

type BotStruct struct {
	bot    *tgbotapi.BotAPI
	logger *slog.Logger
}

func NewTgBotInitialize(bot *tgbotapi.BotAPI, lg *slog.Logger) *BotStruct {
	return &BotStruct{bot: bot, logger: lg}
}

func (b *BotStruct) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		b.logger.Error("GetUpdatesChan error", slog.String("error", err.Error()))
		return err
	}
	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			b.bot.Send(msg)
		}
	}
	return nil
}
