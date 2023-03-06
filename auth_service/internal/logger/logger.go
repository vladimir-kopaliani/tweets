package logger

import (
	"context"
	"os"

	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"

	"github.com/rs/zerolog"
)

func NewLogger(conf Configuration) interfaces.Logger {
	const skipFrameCount = 3
	var l *Logger

	if conf.IsDebugMode {
		// human-friendly output
		l = &Logger{
			stdoutLogger: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
				With().
				Timestamp().
				CallerWithSkipFrameCount(skipFrameCount).
				Logger(),
			stderrLogger: zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
				With().
				Timestamp().
				CallerWithSkipFrameCount(skipFrameCount).
				Logger(),
		}
	} else {
		l = &Logger{
			stdoutLogger: zerolog.New(os.Stdout).
				With().
				Timestamp().
				CallerWithSkipFrameCount(skipFrameCount).
				Logger(),
			stderrLogger: zerolog.New(os.Stderr).
				With().
				Timestamp().
				CallerWithSkipFrameCount(skipFrameCount).
				Logger(),
		}

		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	}

	return l
}

type Logger struct {
	stdoutLogger zerolog.Logger
	stderrLogger zerolog.Logger
}

func (l Logger) Error(ctx context.Context, msg string) {
	l.stderrLogger.Error().Msg(msg)
}

func (l Logger) Warn(ctx context.Context, msg string) {
	l.stdoutLogger.Warn().Msg(msg)
}

func (l Logger) Info(ctx context.Context, msg string) {
	l.stdoutLogger.Info().Msg(msg)
}

func (l Logger) Debug(ctx context.Context, msg string) {
	l.stdoutLogger.Debug().Msg(msg)
}
