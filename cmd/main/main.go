package main

import (
	"context"
	"errors"
	telegram "github.com/Sanchir01/sandjma_graphql/internal/bot"
	"github.com/Sanchir01/sandjma_graphql/internal/config"
	colorStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/color"
	"github.com/Sanchir01/sandjma_graphql/internal/server/grpc/authgrpc"
	"github.com/Sanchir01/sandjma_graphql/internal/server/grpc/categorygrpc"
	"github.com/Sanchir01/sandjma_graphql/pkg/lib/db/connect"

	"github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	sizeStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/size"
	userStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/user"
	httpHandlers "github.com/Sanchir01/sandjma_graphql/internal/handlers"
	httpServer "github.com/Sanchir01/sandjma_graphql/internal/server/http"
	"github.com/Sanchir01/sandjma_graphql/pkg/lib/logger/handlers/slogpretty"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/go-chi/chi/v5"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

	db := connect.PostgresCon(cfg, lg)

	lg.Info("DATABASE CONNECTED", db)
	defer db.Close()
	trManager := manager.Must(trmsqlx.NewDefaultFactory(db))
	r := chi.NewRouter()

	authgrpclient, err := authgrpc.NewGrpcAuth(context.Background(), cfg.GrpcClients.Auth.Address, cfg.GrpcClients.Auth.Retries, lg)

	if err != nil {
		lg.Error("failed init auth grpc client", err)
		os.Exit(1)
	}
	categorygrpcclient, err := categorygrpc.NewGrpcCategory(context.Background(), cfg.GrpcClients.Category.Address, cfg.GrpcClients.Category.Retries, lg)
	if err != nil {
		lg.Error("failed init category grpc client", err)
		os.Exit(1)
	}
	var (
		productStorage = productStore.NewProductPostgresStorage(db)
		userStorages   = userStorage.NewUserPostgresStorage(db)
		colorStorages  = colorStorage.NewColorPostgresStorage(db)
		sizeStorages   = sizeStorage.NewProductPostgresStorage(db)
		handlers       = httpHandlers.NewChiRouter(
			lg, cfg, r,
			productStorage, categorygrpcclient, userStorages, sizeStorages, colorStorages,
			trManager, authgrpclient,
		)
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
		if err := serve.Run(handlers.StartHttpServer()); err != nil {
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
	var lg *slog.Logger
	switch env {
	case development:
		lg = setupPrettySlog()
	case production:
		lg = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return lg
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
