package customMiddleware

import (
	"context"
	"errors"
	userFeature "github.com/Sanchir01/sandjma_graphql/internal/feature/user"
	"net/http"
)

const responseWriterKey = "responseWriter"

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			access, err := r.Cookie("accessToken")
			if err != nil {
				refresh, err := r.Cookie("refreshToken")
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				accessToken, err := userFeature.NewAccessToken(refresh.Value, 0, w)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				token, err := userFeature.ParseToken(accessToken)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}
				ctx := context.WithValue(r.Context(), "user", token)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			_, err = userFeature.ParseToken(access.Value)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// Проверка валидности токена
			validAccessToken, err := userFeature.ParseToken(access.Value)
			if err != nil {

				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), "user", validAccessToken)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetJWTClaimsFromCtx(ctx context.Context) (*userFeature.Claims, error) {
	claims, ok := ctx.Value("user").(*userFeature.Claims)
	if !ok {
		return nil, errors.New("no JWT claims found in context")
	}
	return claims, nil
}

func WithResponseWriter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), responseWriterKey, w)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func GetResponseWriter(ctx context.Context) http.ResponseWriter {
	return ctx.Value(responseWriterKey).(http.ResponseWriter)
}
