package httpHandlers

import (
	"context"
	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sanchir01/sandjma_graphql/internal/config"
	categoryStore "github.com/Sanchir01/sandjma_graphql/internal/database/store/category"
	productStore "github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	storage "github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	userStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/user"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/directive"
	genGql "github.com/Sanchir01/sandjma_graphql/internal/gql/generated"
	resolver "github.com/Sanchir01/sandjma_graphql/internal/gql/resolvers"
	customMiddleware "github.com/Sanchir01/sandjma_graphql/internal/handlers/middleware"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"log/slog"
	"net/http"
	"runtime"
)

type Router struct {
	chiRouter   *chi.Mux
	lg          *slog.Logger
	config      *config.Config
	productStr  *productStore.ProductPostgresStorage
	categoryStr *categoryStore.CategoryPostgresStore
	userStr     *userStorage.UserPostgresStorage
	db          *sqlx.DB
	trManager   *manager.Manager
}

const (
	maxUploadSize                       = 30 * 1024 * 1024
	queryCacheLRUSize                   = 1000
	complexityLimit                     = 1000
	automaticPersistedQueryCacheLRUSize = 100
)

func NewChiRouter(
	lg *slog.Logger, config *config.Config, chi *chi.Mux, productStr *storage.ProductPostgresStorage,
	categoryStr *categoryStore.CategoryPostgresStore, userStr *userStorage.UserPostgresStorage, db *sqlx.DB,
	trManager *manager.Manager) *Router {
	return &Router{
		chiRouter:   chi,
		lg:          lg,
		config:      config,
		productStr:  productStr,
		categoryStr: categoryStr,
		userStr:     userStr,
		db:          db,
		trManager:   trManager,
	}
}

func (rout *Router) StartHttpServer() http.Handler {
	rout.newChiCors()
	rout.chiRouter.Use(middleware.RequestID)
	rout.chiRouter.Use(customMiddleware.WithResponseWriter)
	rout.chiRouter.Use(customMiddleware.AuthMiddleware())

	rout.chiRouter.Handle("/graphql", playground.ApolloSandboxHandler("Sandjma", "/"))
	rout.chiRouter.Handle("/", rout.NewGraphQLHandler())

	return rout.chiRouter
}

func (rout *Router) NewGraphQLHandler() *gqlhandler.Server {

	srv := gqlhandler.New(
		genGql.NewExecutableSchema(rout.newSchemaConfig()),
	)
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.SetQueryCache(lru.New(queryCacheLRUSize))
	srv.AddTransport(transport.MultipartForm{
		MaxUploadSize: maxUploadSize,
		MaxMemory:     maxUploadSize / 10,
	})
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{Cache: lru.New(automaticPersistedQueryCacheLRUSize)})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		buf := make([]byte, 1024)
		n := runtime.Stack(buf, false)
		log.Printf("Panic: %v\nStack: %s\n", err, buf[:n])

		return gqlerror.Errorf("internal server error graphql обработка паники")
	})
	srv.Use(extension.FixedComplexityLimit(complexityLimit))

	return srv
}

func (rout *Router) newSchemaConfig() genGql.Config {
	cfg := genGql.Config{Resolvers: resolver.NewResolver(rout.productStr, rout.categoryStr, rout.userStr, rout.lg, rout.db, rout.trManager)}
	cfg.Directives.InputUnion = directive.NewInputUnionDirective()
	cfg.Directives.SortRankInput = directive.NewSortRankInputDirective()
	cfg.Directives.HasRole = directive.RoleDirective()

	return cfg
}

func (rout *Router) newChiCors() {
	rout.chiRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:           300,
	}))
}
