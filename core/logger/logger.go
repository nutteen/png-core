package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitializeLogger(config LoggerConfig)  {
	if config.IsProductionMode {
		Log, _ = zap.NewProduction()
	} else {
		Log, _ = zap.NewDevelopment()
	}
	defer Log.Sync()
}

func LogContext(ctxProvider LogContextProvider) zap.Field {
	return zap.Inline(contextProvider{ctxProvider})
}

type LogContextProvider interface {
	Value(key interface{}) interface{}
}

type (
	contextProvider struct {
		LogContextProvider
	}
)

func (v contextProvider) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("request_id", v.Value("requestid").(string))
	return nil
}


// LoggerConfig defines the config for middleware.
type LoggerConfig struct {
	IsProductionMode bool
}