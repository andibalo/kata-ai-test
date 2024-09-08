package logger

import (
	"github.com/rs/zerolog"
	"os"
	"sync"
)

var (
	initLoggerOnce sync.Once
	logger         *zerolog.Logger
)

func GetLogger() *zerolog.Logger {
	initLoggerOnce.Do(func() {
		logger = InitLogger()

	})

	return logger
}

type RequestHook struct{}

func (h RequestHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()

	requestID := ctx.Value("request_id")

	if requestID != nil {
		e.Str("request_id", requestID.(string))
	}
}

func InitLogger() *zerolog.Logger {

	l := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger()

	var hooks []zerolog.Hook

	hooks = append(hooks, RequestHook{})

	l = l.Hook(hooks...)

	return &l
}
