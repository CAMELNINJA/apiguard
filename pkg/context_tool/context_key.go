package contexttool

import "context"

type ContextKey int

const (
	// RequestIDKey is the key for the request ID in the context
	RequestIDKey ContextKey = iota
	// UserIDKey is the key for the user ID in the context
	UserIDKey
	// AuthTokenKey is the key for the auth token in the context
	AuthTokenKey
	// APIKeyKey is the key for the API key in the context
	APIKeyKey
	// RateLimitKey is the key for the rate limit in the context
	RateLimitKey
	// TraceIDKey is the key for the trace ID in the context
	TraceIDKey
)

func (k ContextKey) Get(ctx context.Context) (any, bool) {
	value := ctx.Value(k)
	if value == nil {
		return nil, false
	}

	return value, true
}

func (k ContextKey) Set(ctx context.Context, value any) context.Context {
	return context.WithValue(ctx, k, value)
}
