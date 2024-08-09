package resolver

import (
	categoryStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/category"
	colorStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/color"
	productStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	sizeStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/size"
	userStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/user"
	"github.com/Sanchir01/sandjma_graphql/internal/server/grpc/authgrpc"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"log/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductStr    *productStorage.ProductPostgresStorage
	CategoryStr   *categoryStorage.CategoryPostgresStore
	UserStr       *userStorage.UserPostgresStorage
	Logger        *slog.Logger
	SizeStr       *sizeStorage.SizePostgresStorage
	ColorStr      *colorStorage.ColorPostgresStorage
	TrManager     *manager.Manager
	Authgrpclient *authgrpc.Client
}

func NewResolver(
	ProductStr *productStorage.ProductPostgresStorage,
	CategoryStr *categoryStorage.CategoryPostgresStore,
	UserStr *userStorage.UserPostgresStorage,
	Logger *slog.Logger,
	SizeStr *sizeStorage.SizePostgresStorage,
	ColorStr *colorStorage.ColorPostgresStorage,
	TrManager *manager.Manager,
	authgrpclient *authgrpc.Client,
) *Resolver {
	return &Resolver{
		ProductStr:    ProductStr,
		CategoryStr:   CategoryStr,
		UserStr:       UserStr,
		Logger:        Logger,
		SizeStr:       SizeStr,
		ColorStr:      ColorStr,
		TrManager:     TrManager,
		Authgrpclient: authgrpclient,
	}
}
