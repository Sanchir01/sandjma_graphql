package categoryStore

import (
	"context"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"log/slog"
	"time"
)

type CategoryPostgresStore struct {
	db *sqlx.DB
}

func NewCategoryPostgresStore(db *sqlx.DB) *CategoryPostgresStore {
	return &CategoryPostgresStore{
		db: db,
	}
}

func (db *CategoryPostgresStore) GetAllCategory(ctx context.Context) ([]model.Category, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var category []dbCategory
	if err := conn.SelectContext(ctx, &category, "SELECT * FROM categories"); err != nil {
		return nil, err
	}
	return lo.Map(category, func(category dbCategory, _ int) model.Category { return model.Category(category) }), nil
}

type dbCategory struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Slug        string    `db:"slug"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Description string    `db:"description"`
	Version     uint      `db:"version"`
}

func (db *CategoryPostgresStore) CreateCategory(ctx context.Context, input *model.CreateCategoryInput, slug string) (uuid.UUID, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return uuid.New(), err
	}
	defer conn.Close()
	var id uuid.UUID
	row := conn.QueryRowContext(ctx,
		"INSERT INTO categories(name, slug, description) VALUES($1, $2, $3) RETURNING id",
		input.Name, slug, input.Description)
	if err := row.Err(); err != nil {
		slog.Error("create category error", err.Error())
		return uuid.New(), err
	}
	if err := row.Scan(&id); err != nil {
		return uuid.New(), err
	}
	return id, nil
}
