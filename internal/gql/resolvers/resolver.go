package resolver

import (
	categoryStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/category"
	productStorage "github.com/Sanchir01/sandjma_graphql/internal/database/store/product"
	"log/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductStr  *productStorage.ProductPostgresStorage
	CategoryStr *categoryStorage.CategoryPostgresStore
	Logger      *slog.Logger
}
