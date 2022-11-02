package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	LogContextProvider interface {
		Value(key interface{}) interface{}
	}

	contextProvider struct {
		LogContextProvider
	}
)

var Log *zap.Logger

func InitializeLogger(logType string) {
	if logType == "prod" {
		Log, _ = zap.NewProduction()
	} else {
		Log, _ = zap.NewDevelopment()
	}
	defer Log.Sync()
}

func LogContext(ctxProvider LogContextProvider) zap.Field {
	return zap.Inline(contextProvider{LogContextProvider: ctxProvider})
}

func (v contextProvider) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("request_id", v.Value("requestid").(string))
	return nil
}
