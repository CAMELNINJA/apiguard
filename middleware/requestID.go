package middleware

import (
	"net/http"

	contexttool "github.com/CAMELNINJA/apiguard/pkg/context_tool"
	uuid "github.com/satori/go.uuid"
)

func RequestId(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get("X-Request-ID")
		if rid == "" {
			rid = uuid.NewV4().String()
		}

		ctx := contexttool.RequestIDKey.Set(r.Context(), rid)
		rw.Header().Add("X-Request-ID", rid)

		next.ServeHTTP(rw, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
