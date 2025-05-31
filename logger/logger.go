package logger

import (
	"encoding/json"

	"github.com/wisaitas/share-pkg/errs"
	"github.com/wisaitas/share-pkg/notify/discord"
	"go.uber.org/zap"
)

type Logger interface {
	Info(message string)
	ErrorWithDiscord(err error)
	ErrorWithoutDiscord(err error)
	Warn(message string)
}

type Options struct {
	Discord discord.Discord
}

type logger struct {
	zlog    *zap.Logger
	discord discord.Discord
}

func NewLogger(zlog *zap.Logger, options *Options) Logger {
	l := &logger{
		zlog: zlog.WithOptions(zap.AddCallerSkip(1)),
	}

	if options != nil && options.Discord != nil {
		l.discord = options.Discord
	}

	return l
}

func (l *logger) Info(message string) {
	l.zlog.Info(message)
}

func (l *logger) ErrorWithDiscord(err error) {
	if err != nil {
		l.zlog.Error(err.Error())

		if l.discord != nil {
			message, err := errs.ErrorMessageWithoutLog(err)
			if err != nil {
				l.zlog.Error("failed to get error message", zap.Error(err))
			}

			jsonMessage, err := json.Marshal(message)
			if err != nil {
				l.zlog.Error("failed to marshal error message", zap.Error(err))
			}

			if discordErr := l.discord.SendMessage(string(jsonMessage)); discordErr != nil {
				l.zlog.Error("failed to send message to discord", zap.Error(discordErr))
			}
		}
	}
}

func (l *logger) ErrorWithoutDiscord(err error) {
	if err != nil {
		l.zlog.Error(err.Error())
	}
}

func (l *logger) Warn(message string) {
	l.zlog.Warn(message)
}
