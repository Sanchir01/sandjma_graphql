package main

import (
	"context"
	"errors"
	telegram "github.com/Sanchir01/sandjma_graphql/internal/bot"
	"github.com/Sanchir01/sandjma_graphql/internal/config"
	storage "github.com/Sanchir01/sandjma_graphql/internal/database/store"
	httpHandlers "github.com/Sanchir01/sandjma_graphql/internal/handlers"
	httpServer "github.com/Sanchir01/sandjma_graphql/internal/server/http"
	"github.com/Sanchir01/sandjma_graphql/pkg/lib/logger/handlers/slogpretty"
	"github.com/go-chi/chi/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var (
	development = "development"
	production  = "production"
)

func main() {

	cfg := config.InitConfig()
	lg := setupLogger(cfg.Env)
	lg.Info("Graphql server starting up...", slog.String("port", cfg.HttpServer.Port))
	db, err := sqlx.Connect("postgres", "user=postgres dbname=golangS sslmode=disable password=sanchirgarik01")
	if err != nil {
		lg.Error("sqlx.Connect error", slog.String("error", err.Error()))
	}
	lg.Info("DATABASE CONNECTED", db)

	defer db.Close()

	r := chi.NewRouter()
	var (
		productStorage = storage.NewProductPostgresStorage(db)
		handlers       = httpHandlers.NewChiRouter(lg, cfg, r, productStorage)
	)
	serve := httpServer.NewHttpServer(cfg)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer cancel()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	telegramBot := telegram.NewTgBotInitialize(bot, lg)
	go func(ctx context.Context) {
		if err := serve.Run(handlers.StartHttpHandlers()); err != nil {
			if !errors.Is(err, context.Canceled) {
				lg.Error("Listen server error", slog.String("error", err.Error()))
				return
			}
			lg.Error("Listen server error", slog.String("error", err.Error()))
		}
	}(ctx)

	if err := telegramBot.Start(ctx); err != nil {
		if !errors.Is(err, context.Canceled) {
			lg.Error("Telegram bot error", slog.String("error", err.Error()))
			return
		}
		lg.Error("Telegram bot stoped", slog.String("error", err.Error()))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case development:
		log = setupPrettySlog()
	case production:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
