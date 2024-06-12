package middleware

import (
	"context"
	"net/http"
)

type injectableResponseWriter struct {
	http.ResponseWriter
	Cookie *http.Cookie
}

func (i *injectableResponseWriter) Write(data []byte) (int, error) {
	if i.Cookie != nil {
		http.SetCookie(i.ResponseWriter, i.Cookie)
	}

	return i.ResponseWriter.Write(data)
}

func SomeMiddleware(someParam string, inner http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := req.Context()

			// pull cookie from request
			cookie, err := req.Cookie("refreshToken")

			// since err can only return ErrNoCookie lets only add cookie to
			// context if the err is nil.
			if err == nil && cookie != nil {
				ctx = context.WithValue(
					ctx,
					"COOKIE_VALUE",
					cookie,
				)
			}

			// add writer pointer for returnable cookies
			cres := injectableResponseWriter{ResponseWriter: res}
			ctx = context.WithValue(ctx, "RESPONSE_WRITER", &cres)

			// pass request on with injected writer and cookie value
			next.ServeHTTP(&cres, req.WithContext(ctx))
		})
	}(inner)
}
