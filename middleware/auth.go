package middleware

import (
	"net/http"
	"strings"

	"github.com/CAMELNINJA/apiguard/config"
	contexttool "github.com/CAMELNINJA/apiguard/pkg/context_tool"
)

func AuthOptional(cfg *config.Config) func(http.Handler) http.Handler {
	if len(cfg.ApiKeys) == 0 {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				next.ServeHTTP(w, r)
			})
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				next.ServeHTTP(w, r)
				return
			}

			token := strings.TrimPrefix(auth, "Bearer ")

			_, ok := cfg.MapApiKeys[token]
			if ok {
				ctx := contexttool.AuthTokenKey.Set(r.Context(), token)
				next.ServeHTTP(w, r.WithContext(ctx))

				return
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}

func AuthRequired() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, ok := contexttool.AuthTokenKey.Get(r.Context())
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
