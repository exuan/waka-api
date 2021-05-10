package api

import (
	"github.com/exuan/waka-api/internal/config"
	"github.com/exuan/waka-api/internal/logger"
	"github.com/exuan/waka-api/service"
	"go.uber.org/zap/zapcore"
)

const (
	ErrorLevel = iota + zapcore.FatalLevel + 1
	RequestLevel
	LogPathFormat = "%Y-%m-%d"
)

var (
	Cfg     *config.Config
	Log     *logger.Sugar
	Service service.Service
	LogCfgs = []*logger.Config{
		{
			Name:   "error",
			Level:  ErrorLevel,
			Format: LogPathFormat,
		},
		{
			Name:   "request",
			Level:  RequestLevel,
			Format: LogPathFormat,
		},
	}
)
