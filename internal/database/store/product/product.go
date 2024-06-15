package productStore

import (
	"context"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"log/slog"
	"time"
)

type ProductPostgresStorage struct {
	db *sqlx.DB
}

func NewProductPostgresStorage(db *sqlx.DB) *ProductPostgresStorage {
	return &ProductPostgresStorage{db: db}
}

func (p *ProductPostgresStorage) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	conn, err := p.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var products []dbProduct
	if err = conn.SelectContext(ctx, &products, "SELECT * FROM products"); err != nil {
		return nil, err
	}
	return lo.Map(products, func(products dbProduct, _ int) model.Product { return model.Product(products) }), nil
}

func (p *ProductPostgresStorage) CreateProduct(ctx context.Context, input *model.CreateProductInput) (uuid.UUID, error) {
	conn, err := p.db.Connx(ctx)
	if err != nil {
		return uuid.New(), err
	}
	defer conn.Close()

	var id uuid.UUID

	row := conn.QueryRowContext(ctx,
		"INSERT INTO products(name, price, category_id, description) VALUES($1, $2, $3, $4) RETURNING id",
		input.Name, input.Price, input.CategoryID, input.Description)

	if err := row.Err(); err != nil {
		slog.Error("create product error", err)
		return uuid.New(), err
	}
	if err := row.Scan(&id); err != nil {
		return uuid.New(), err
	}
	return id, nil

}

type dbProduct struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Price       int       `db:"price"`
	Images      []string  `db:"images"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	CategoryID  uuid.UUID `db:"category_id"`
	Description string    `db:"description"`
	Version     uint      `db:"version"`
}
