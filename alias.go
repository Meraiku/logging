package logging

import (
	"log/slog"
	"time"
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

type (
	Logger         = slog.Logger
	Level          = slog.Level
	Attr           = slog.Attr
	Value          = slog.Value
	LogValuer      = slog.LogValuer
	Handler        = slog.Handler
	HandlerOptions = slog.HandlerOptions
)

var (
	New            = slog.New
	NewTextHandler = slog.NewTextHandler
	NewJSONHandler = slog.NewJSONHandler
	SetDefault     = slog.SetDefault

	String   = slog.String
	Bool     = slog.Bool
	Float64  = slog.Float64
	Any      = slog.Any
	Duration = slog.Duration
	Int      = slog.Int
	Int64    = slog.Int64
	Uint64   = slog.Uint64

	Group      = slog.Group
	GroupValue = slog.GroupValue
)

func Time(key string, t time.Time) Attr {
	return String(key, t.String())
}

func Err(err error) Attr {
	return String("error", err.Error())
}
