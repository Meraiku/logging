package logging

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"

	slogmulti "github.com/samber/slog-multi"
)

const (
	defaultLevel      = LevelInfo
	defaultIsJSON     = true
	defaultAddSource  = true
	defaultSetDefault = true
	defaultLogstash   = false
)

func NewLogger(opts ...LoggerOption) *Logger {
	cfg := LoggerOptions{
		Level:      defaultLevel,
		IsJSON:     defaultIsJSON,
		AddSource:  defaultAddSource,
		SetDefault: defaultSetDefault,
		Logstash: Logstash{
			Enable: defaultLogstash,
		},
	}

	for _, opt := range opts {
		opt(&cfg)
	}

	ho := &HandlerOptions{
		Level:     cfg.Level,
		AddSource: cfg.AddSource,
	}

	var h Handler

	switch cfg.IsJSON {
	case true:
		h = NewJSONHandler(os.Stdout, ho)
	case false:
		h = NewTextHandler(os.Stdout, ho)
	}

	if cfg.Logstash.Enable {
		conn, err := net.Dial("udp", cfg.Logstash.Addr)
		if err != nil {
			log.Fatalf("failed to connect to logstash: %v", err)
		}
		h = slogmulti.Fanout(h, NewJSONHandler(conn, ho))
	}

	l := New(h)

	if cfg.SetDefault {
		SetDefault(l)
	}

	return l
}

type LoggerOptions struct {
	Level      Level
	IsJSON     bool
	AddSource  bool
	SetDefault bool
	Logstash   Logstash
}

type Logstash struct {
	Enable bool
	Addr   string
}

type LoggerOption func(*LoggerOptions)

// WithLevel sets the logging level for the logger. Default is LevelInfo.
func WithLevel(level string) LoggerOption {
	return func(o *LoggerOptions) {
		var l Level

		if err := l.UnmarshalText([]byte(level)); err != nil {
			l = LevelInfo
		}

		o.Level = l
	}
}

// WithJSON sets whether the logger should output JSON-formatted logs.
func WithJSON(isJSON bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.IsJSON = isJSON
	}
}

// WithSource sets whether the logger should include source information.
func WithSource(addSource bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.AddSource = addSource
	}
}

// WithSetDefault sets whether the logger should be set as the default logger.
func WithSetDefault(setDefault bool) LoggerOption {
	return func(o *LoggerOptions) {
		o.SetDefault = setDefault
	}
}

// WithLogstash sets whether the logger should send logs to Logstash.
func WithLogstash(enable bool, logstashAddress string) LoggerOption {
	return func(o *LoggerOptions) {
		logstash := Logstash{
			Enable: enable,
			Addr:   logstashAddress,
		}
		o.Logstash = logstash
	}
}

// WithAttrs sets attributes for the logger.
func WithAttrs(ctx context.Context, attrs ...Attr) *Logger {
	logger := L(ctx)

	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

// WithDefaultAttrs sets default attributes for the logger.
func WithDefaultAttrs(logger *Logger, attrs ...Attr) *Logger {

	for _, attr := range attrs {
		logger = logger.With(attr)
	}

	return logger
}

// L returns a logger from the context.
func L(ctx context.Context) *Logger {
	return loggerFromContext(ctx)
}

// Default returns the default logger.
func Default() *Logger {
	return slog.Default()
}
