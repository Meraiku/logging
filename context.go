package logging

import "context"

type key struct{}

func ContextWithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, key{}, l)
}

func loggerFromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(key{}).(*Logger); ok {
		return l
	}

	return Default()
}
