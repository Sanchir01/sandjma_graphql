package httpHandlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Sanchir01/sandjma_graphql/internal/config"
	categoryStore "github.com/Sanchir01/sandjma_graphql/internal/database/store/category"
	productStore "github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	storage "github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	userStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/user"
	runtime "github.com/Sanchir01/sandjma_graphql/internal/gql/generated"
	resolver "github.com/Sanchir01/sandjma_graphql/internal/gql/resolvers"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

type Router struct {
	chiRouter   *chi.Mux
	lg          *slog.Logger
	config      *config.Config
	productStr  *productStore.ProductPostgresStorage
	categoryStr *categoryStore.CategoryPostgresStore
	userStr     *userStorage.UserPostgresStorage
}

func NewChiRouter(
	lg *slog.Logger, config *config.Config, chi *chi.Mux, productStr *storage.ProductPostgresStorage,
	categoryStr *categoryStore.CategoryPostgresStore, userStr *userStorage.UserPostgresStorage) *Router {
	return &Router{
		chiRouter:   chi,
		lg:          lg,
		config:      config,
		productStr:  productStr,
		categoryStr: categoryStr,
		userStr:     userStr,
	}
}

func (rout *Router) StartHttpHandlers() http.Handler {
	srv := handler.NewDefaultServer(runtime.NewExecutableSchema(runtime.Config{Resolvers: &resolver.Resolver{
		UserStr:     rout.userStr,
		ProductStr:  rout.productStr,
		CategoryStr: rout.categoryStr,
		Logger:      rout.lg,
	}}))

	rout.chiRouter.Handle("/", playground.ApolloSandboxHandler("GraphQL playground", "/query"))
	rout.chiRouter.Handle("/query", srv)
	return rout.chiRouter
}
