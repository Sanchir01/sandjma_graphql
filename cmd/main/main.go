package main

import (
	"context"
	"errors"
	telegram "github.com/Sanchir01/sandjma_graphql/internal/bot"
	"github.com/Sanchir01/sandjma_graphql/internal/config"
	categoryStore "github.com/Sanchir01/sandjma_graphql/internal/database/store/category"
	"github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	userStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/user"
	httpHandlers "github.com/Sanchir01/sandjma_graphql/internal/handlers"
	httpServer "github.com/Sanchir01/sandjma_graphql/internal/server/http"
	"github.com/Sanchir01/sandjma_graphql/pkg/lib/logger/handlers/slogpretty"
	trmsqlx "github.com/avito-tech/go-transaction-manager/drivers/sqlx/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
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

type ctxKey int

var (
	development = "development"
	production  = "production"
)

func main() {

	cfg := config.InitConfig()
	lg := setupLogger(cfg.Env)
	lg.Info("Graphql server starting up...", slog.String("port", cfg.HttpServer.Port))
	db, err := sqlx.Open("postgres", "user=postgres dbname=golangS sslmode=disable password=sanchirgarik01")
	if err != nil {
		lg.Error("sqlx.Connect error", slog.String("error", err.Error()))
	}
	lg.Info("DATABASE CONNECTED", db)

	defer db.Close()
	trManager := manager.Must(trmsqlx.NewDefaultFactory(db))
	r := chi.NewRouter()

	var (
		productStorage  = productStore.NewProductPostgresStorage(db)
		categoryStorage = categoryStore.NewCategoryPostgresStore(db)
		userStorages    = userStorage.NewUserPostgresStorage(db)
		handlers        = httpHandlers.NewChiRouter(lg, cfg, r, productStorage, categoryStorage, userStorages, db, trManager)
	)
	//id, err := productStorage.CreateProduct(context.Background(), &model.CreateProductInput{
	//	Name:        "just",
	//	Description: "test",
	//	Price:       123,
	//	CategoryID:  uuid.MustParse("3445eef1-4db1-4bcf-b8e5-c7a46d8b2443"),
	//	Images: []string{
	//		"https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.vogue.ru%2Ffashion%2F7-idej-verhnej-odezhdy-dlya-holodnyh-dnej&psig=AOvVaw0MgOiFOmhl84KWrrlYzjF8&ust=1718541447417000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCMiDzvXP3YYDFQAAAAAdAAAAABAE",
	//		"https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.xn--e1agsbgdeid7b.xn--p1ai%2Farticles%2Fuhod-za-odezhdoj%2Fkak-ubrat-staticheskoe-elektrichestvo-s-odezhdy%2F&psig=AOvVaw0MgOiFOmhl84KWrrlYzjF8&ust=1718541447417000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCMiDzvXP3YYDFQAAAAAdAAAAABAI",
	//	},
	//},
	//)
	//if err != nil {
	//	lg.Error("productStorage.CreateProduct error", slog.String("error", err.Error()))
	//}
	//lg.Warn("ID", slog.String("id", id.String()))
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
