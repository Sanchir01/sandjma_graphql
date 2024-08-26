package colorStorage

import (
	"context"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
	"log/slog"
	"time"
)

type ColorPostgresStorage struct {
	db *sqlx.DB
}

func NewColorPostgresStorage(db *sqlx.DB) *ColorPostgresStorage {
	return &ColorPostgresStorage{db: db}
}

func (db *ColorPostgresStorage) CreateColor(ctx context.Context, input *model.CreateColorInput, slug string) (uuid.UUID, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()
	var id uuid.UUID
	row := conn.QueryRowContext(ctx, "INSERT INTO colors(name, slug, css_variables) VALUES($1, $2, $3) RETURNING id",
		input.Name, slug, input.CSSVariables)
	if err := row.Err(); err != nil {
		slog.Error("create category error", err.Error())
		return uuid.Nil, err
	}
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (db *ColorPostgresStorage) GetAllColor(ctx context.Context) ([]model.Color, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var color []dbColor
	if err := conn.SelectContext(ctx, &color, "SELECT * FROM colors"); err != nil {
		return nil, err
	}
	return lo.Map(color, func(color dbColor, _ int) model.Color { return model.Color(color) }), nil
}

type dbColor struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Slug         string    `db:"slug"`
	CSSVariables string    `db:"css_variables"`
	Version      uint      `db:"version"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
