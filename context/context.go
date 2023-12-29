package context

import (
	"context"
	"time"
)

func New(args *Args) context.Context {
	ctx := context.WithValue(context.Background(), "args", args)
	return ctx
}

func WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}
