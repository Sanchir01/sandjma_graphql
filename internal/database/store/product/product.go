package productStore

import (
	"context"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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

	newProducts := make([]model.Product, len(products))

	for i, dbProd := range products {
		newProducts[i] = model.Product{
			ID:          dbProd.ID,
			Name:        dbProd.Name,
			Price:       dbProd.Price,
			Images:      dbProd.Images,
			CreatedAt:   dbProd.CreatedAt,
			UpdatedAt:   dbProd.UpdatedAt,
			CategoryID:  dbProd.CategoryID,
			Description: dbProd.Description,
			Version:     dbProd.Version,
		}
	}

	return newProducts, nil
}

func (p *ProductPostgresStorage) CreateProduct(ctx context.Context, input *model.CreateProductInput) (uuid.UUID, error) {
	conn, err := p.db.Connx(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()
	myImages := make([]string, len(input.Images))
	copy(myImages, input.Images)
	imagesArray := pq.Array(myImages)
	var id uuid.UUID

	if imagesArray == nil {
		slog.Error("create images error", err.Error())
		return uuid.Nil, err
	}
	row := conn.QueryRowContext(ctx,
		"INSERT INTO products(name, price, category_id, description, images) VALUES($1, $2, $3, $4, $5) RETURNING id",
		input.Name, input.Price, input.CategoryID, input.Description, imagesArray)

	if err := row.Err(); err != nil {
		slog.Error("create product error", err.Error())
		return uuid.Nil, err
	}
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil

}

type dbProduct struct {
	ID          uuid.UUID      `db:"id"`
	Name        string         `db:"name"`
	Price       int            `db:"price"`
	Images      pq.StringArray `db:"images"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
	CategoryID  uuid.UUID      `db:"category_id"`
	Description string         `db:"description"`
	Version     uint           `db:"version"`
}
