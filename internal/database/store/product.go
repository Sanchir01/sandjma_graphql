package store

import (
	"context"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
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

func (p *ProductPostgresStorage) CreateProduct(ctx context.Context, input *model.CreateProductInput) (int32, error) {
	conn, err := p.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int32
	row := conn.QueryRowContext(ctx,
		"INSERT INTO products(name, price, category_id) VALUES($1, $2, $3) RETURNING id",
		input.Name, input.Price, input.CategoryID)
	if err := row.Err(); err != nil {
		return 0, err
	}
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil

}

type dbProduct struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Price      int       `db:"price"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	CategoryID uuid.UUID `db:"category_id"`
	Version    uint      `db:"version"`
}
