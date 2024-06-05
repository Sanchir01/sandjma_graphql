package userStorage

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type UserPostgresStorage struct {
	db *sqlx.DB
}

func NewUserPostgresStorage(db *sqlx.DB) *UserPostgresStorage {
	return &UserPostgresStorage{db: db}
}

func (db *UserPostgresStorage) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var user dbUser

	if err = conn.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return nil, err
	}

	return &model.User{
		ID:         user.ID,
		Name:       user.Name,
		Phone:      user.Phone,
		Email:      user.Email,
		Role:       model.Role(user.Role),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		AvatarPath: user.AvatarPath, // Путь б
	}, nil
}

type dbUser struct {
	ID         uuid.UUID      `db:"id"`
	Name       string         `db:"name"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
	Phone      string         `db:"phone"`
	Email      string         `db:"email"`
	AvatarPath graphql.Upload `db:"avatar_path"`
	Role       Role           `db:"role"`
}

type Role string
