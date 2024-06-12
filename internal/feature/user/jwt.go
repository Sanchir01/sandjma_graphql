package userFeature

import (
	"github.com/Sanchir01/sandjma_graphql/internal/gql/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Claims struct {
	Id   uuid.UUID  `json:"id"`
	Role model.Role `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(id uuid.UUID, Role model.Role, expire time.Time) (string, error) {
	claim := &Claims{
		Id:   id,
		Role: Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
		},
	}

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := tokens.SignedString(secretKey)

	if err != nil {
		slog.Error("error generate jwt token", err.Error())
		return "", err
	}
	return tokenString, nil
}

func AddCookieTokens(id uuid.UUID, Role model.Role) error {
	expirationTimeAccess := time.Now().Add(15 * time.Minute)
	expirationTimeRefresh := time.Now().Add(24 * time.Hour)
	refreshToken, err := GenerateJwtToken(id, Role, expirationTimeRefresh)
	if err != nil {
		return err
	}
	accessToken, err := GenerateJwtToken(id, Role, expirationTimeAccess)
	if err != nil {
		return err
	}
	slog.Warn("tokens generated", slog.String("access_token", accessToken), slog.String("refresh_token", refreshToken))

	return nil
}

func GenerateCookie(name string, expire time.Time, httpOnly bool, value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expire,
		Path:     "/",
		HttpOnly: httpOnly,
	}
	
	return cookie
}
