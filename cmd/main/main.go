package main

import (
	"context"
	"errors"
	"github.com/Sanchir01/sandjma_graphql/internal/config"
	httpHandlers "github.com/Sanchir01/sandjma_graphql/internal/handlers"
	httpServer "github.com/Sanchir01/sandjma_graphql/internal/server/http"
	"github.com/Sanchir01/sandjma_graphql/pkg/lib/logger/handlers/slogpretty"
	"github.com/go-chi/chi/v5"
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
	//dsn := fmt.Sprintf("user=%s password=%s dbname=%s ",
	//	cfg.DB.User,
	//	os.Getenv("DB_PASSWORD"),
	//	cfg.DB.Database,
	//)
	//db, err := sqlx.Connect("postgres", "user=postgres dbname=golangS sslmode=disable password=sanchirgarik01")
	//if err != nil {
	//	lg.Error("sqlx.Connect error", err.Error())
	//}
	//lg.Info("DATABASE CONNECTED", db)
	//
	//defer db.Close()

	r := chi.NewRouter()
	var (
		handlers = httpHandlers.NewChiRouter(lg, cfg, r)
	)
	serve := httpServer.NewHttpServer(cfg)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer cancel()

	go func(ctx context.Context) {
		if err := serve.Run(handlers.StartHttpHandlers()); err != nil {
			if !errors.Is(err, context.Canceled) {
				lg.Error("Listen server error", slog.String("error", err.Error()))
				return
			}
			lg.Error("Listen server error", slog.String("error", err.Error()))
		}
	}(ctx)

	<-ctx.Done()
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
