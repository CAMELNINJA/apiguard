package middleware

import (
	"net/http"
	"time"

	contexttool "github.com/CAMELNINJA/apiguard/pkg/context_tool"
	zap_helper "github.com/CAMELNINJA/apiguard/pkg/zap_once"
	"github.com/go-chi/chi/middleware"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func Logger() func(http.Handler) http.Handler {
	logger := zap_helper.GetLogger()
	return func(next http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(rw, r.ProtoMajor)
			start := time.Now()
			rid, ok := contexttool.RequestIDKey.Get(r.Context())
			if !ok {
				rid = "unknown"
			}

			defer func() {
				logger.Info("request completedt",
					zap.String("request-id", rid.(string)),
					zap.Int("status", ww.Status()),
					zap.Int("bytes", ww.BytesWritten()),
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("query", r.URL.RawQuery),
					zap.String("ip", r.RemoteAddr),
					zap.String("trace.id", trace.SpanFromContext(r.Context()).SpanContext().TraceID().String()),
					zap.String("user-agent", r.UserAgent()),
					zap.Duration("latency", time.Since(start)),
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
