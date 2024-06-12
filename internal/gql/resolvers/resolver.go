package resolver

import (
	categoryStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/category"
	productStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	userStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/user"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductStr  *productStorage.ProductPostgresStorage
	CategoryStr *categoryStorage.CategoryPostgresStore
	UserStr     *userStorage.UserPostgresStorage
	Logger      *slog.Logger
	DB          *sqlx.DB
	TrManager   *manager.Manager
}

func NewResolver(
	ProductStr *productStorage.ProductPostgresStorage,
	CategoryStr *categoryStorage.CategoryPostgresStore,
	UserStr *userStorage.UserPostgresStorage,
	Logger *slog.Logger,
	DB *sqlx.DB,
	TrManager *manager.Manager,
) *Resolver {
	return &Resolver{
		ProductStr:  ProductStr,
		CategoryStr: CategoryStr,
		UserStr:     UserStr,
		Logger:      Logger,
		DB:          DB,
		TrManager:   TrManager,
	}
}
