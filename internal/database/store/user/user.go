package userStorage

import (
	"context"
	userFeature "github.com/Sanchir01/sandjma_graphql/internal/feature/user"
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

	if err := conn.GetContext(ctx, &user, "SELECT * FROM users WHERE email = $1", email); err != nil {
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

func (db *UserPostgresStorage) CreateUser(ctx context.Context, input *model.RegistrationsInput) (*model.User, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var user dbUser
	newPassword, err := userFeature.HashPassword(input.Password)
	err = conn.QueryRowContext(ctx,
		"INSERT INTO users (name, phone, email, role, password) VALUES ($1, $2, $3, $4, $5) RETURNING users.*",
		input.Name, input.Phone, input.Email, input.Role, newPassword).Scan(&user.ID, &user.Name, &user.CreatedAt, &user.UpdatedAt, &user.Phone, &user.Email, &user.AvatarPath, &user.Role, &user.Password)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:         user.ID,
		Name:       user.Name,
		Phone:      user.Phone,
		Password:   user.Password,
		Email:      user.Email,
		Role:       model.Role(user.Role),
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		AvatarPath: user.AvatarPath,
	}, nil
}

func (db *UserPostgresStorage) GetUserByPhone(ctx context.Context, phone string) (*model.User, error) {
	conn, err := db.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	var user dbUser
	if err := conn.GetContext(ctx, &user, "SELECT * FROM users WHERE phone = $1", phone); err != nil {
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
	ID         uuid.UUID `db:"id"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Password   string    `db:"password"`
	Phone      string    `db:"phone"`
	Email      string    `db:"email"`
	AvatarPath string    `db:"avatar_path"`
	Role       Role      `db:"role"`
}

type Role string
