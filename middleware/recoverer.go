package middleware

import (
	"fmt"
	"net/http"

	contexttool "github.com/CAMELNINJA/apiguard/pkg/context_tool"
	zap_helper "github.com/CAMELNINJA/apiguard/pkg/zap_once"
	"github.com/go-stack/stack"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func Recover() func(http.Handler) http.Handler {
	logger := zap_helper.GetLogger()

	return func(next http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, r *http.Request) {
			defer func() {
				if p := recover(); p != nil {
					err, ok := p.(error)
					if !ok {
						err = fmt.Errorf("%v", p)
					}

					var stackTrace stack.CallStack
					// Get the current stacktrace but trim the runtime
					traces := stack.Trace().TrimRuntime()

					// Format the stack trace removing the clutter from it
					for i := 0; i < len(traces); i++ {
						t := traces[i]
						tFunc := t.Frame().Function

						// Opentelemetry is recovering from the panics on span.End defets and throwing them again
						// we don't want this noise to appear on our logs
						if tFunc == "runtime.gopanic" || tFunc == "go.opentelemetry.io/otel/sdk/trace.(*span).End" {
							continue
						}

						// This call is made before the code reaching our handlers, we don't want to log things that are coming before
						// our own code, just from our handlers and donwards.
						if tFunc == "net/http.HandlerFunc.ServeHTTP" {
							break
						}
						stackTrace = append(stackTrace, t)
					}

					reqID, ok := contexttool.RequestIDKey.Get(r.Context())
					if !ok {
						reqID = "unknown"
					}

					fields := []zap.Field{
						zap.Error(err),
						zap.String("trace_id", trace.SpanFromContext(r.Context()).SpanContext().TraceID().String()),
						zap.String("request_id", reqID.(string)),
						zap.String("stack", fmt.Sprintf("%+v", stackTrace)),
						zap.String("method", r.Method),
						zap.String("path", r.URL.Path),
						zap.String("remote_addr", r.RemoteAddr),
						zap.String("user_agent", r.UserAgent()),
						zap.String("request_id", r.Header.Get("X-Request-ID")),
					}

					logger.Panic("panic", fields...)

					http.Error(rw, http.StatusText(http.StatusInternalServerError),
						http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(rw, r)
		}
		return http.HandlerFunc(fn)
	}
}
