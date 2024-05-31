package resolver

import (
	storage "github.com/Sanchir01/sandjma_graphql/internal/database/store"
	"log/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ProductStr *storage.ProductPostgresStorage
	Logger     *slog.Logger
}
