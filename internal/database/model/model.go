package modelDB

import (
	"github.com/google/uuid"
	"time"
)

type ProductDB struct {
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	Price      int       `db:"price"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	CategoryID uuid.UUID `db:"category_id"`
	Version    uint      `db:"version"`
}

type Categories struct {
	Id          int32
	Name        string
	Slug        string
	Description string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
