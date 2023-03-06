package interfaces

import "context"

type Logger interface {
	Error(ctx context.Context, msg string)
	Warn(ctx context.Context, msg string)
	Info(ctx context.Context, msg string)
	Debug(ctx context.Context, msg string)
}
