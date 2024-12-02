package logging

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContextWithLogger(t *testing.T) {
	ctx := context.Background()

	log := NewLogger()

	ctxWithLogger := ContextWithLogger(ctx, log)

	ctxLogger, ok := ctxWithLogger.Value(key{}).(*Logger)
	require.True(t, ok, "expected logger in context")
	require.Equal(t, log, ctxLogger, "logger from context does not match")
}

func TestLoggerFromContext(t *testing.T) {
	ctx := context.Background()

	log := NewLogger()

	ctxWithLogger := ContextWithLogger(ctx, log)

	ctxLogger := loggerFromContext(ctxWithLogger)
	require.Equal(t, log, ctxLogger, "logger from context does not match")
}

func TestLoggerFromContext_WithNoLogger(t *testing.T) {
	ctx := context.Background()

	ctxLogger := loggerFromContext(ctx)
	require.Equal(t, Default(), ctxLogger, "Did not retrive default logger from context")
}
