package sizeStorage

import (
	"context"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"time"
)

type SizePostgresStorage struct {
	db *sqlx.DB
}

func NewProductPostgresStorage(db *sqlx.DB) *SizePostgresStorage {
	return &SizePostgresStorage{db: db}
}

func (db *SizePostgresStorage) GetAllSizes(ctx context.Context) ([]model.Size, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var sizes []dbSize
	if err := conn.SelectContext(ctx, &sizes, "SELECT * FROM size"); err != nil {
		return nil, err
	}

	return lo.Map(sizes, func(size dbSize, _ int) model.Size { return model.Size(size) }), nil
}

func (db *SizePostgresStorage) CreateSize(ctx context.Context, input *model.SizeCreateInput, slug string) (uuid.UUID, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()
	var id uuid.UUID

	row := conn.QueryRowContext(ctx, "INSERT INTO size(name,slug) VALUES($1, $2) RETURNING id", input.Name, slug)

	if err := row.Err(); err != nil {
		return uuid.Nil, err
	}
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

type dbSize struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Slug      string    `db:"slug"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Version   uint      `db:"version"`
}
