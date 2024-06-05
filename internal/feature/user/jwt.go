package userFeature

import (
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"time"
)

type Claims struct {
	Id   uuid.UUID  `json:"username"`
	Role model.Role `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(id uuid.UUID, Role model.Role) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claim := &Claims{
		Id:   id,
		Role: Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := tokens.SignedString(os.Getenv("JWT_SECRET"))

	if err != nil {
		slog.Error("error generate jwt token", err.Error())
		return "", nil
	}
	return tokenString, nil
}
