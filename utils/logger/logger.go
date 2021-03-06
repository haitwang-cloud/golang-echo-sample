package logger

import (
	"errors"
	"fmt"
	"github.com/haitwang-cloud/golang-echo-sample/utils/config"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"

	"go.uber.org/zap"
)

// Logger is an alternative implementation of *gorm.Logger
type Logger struct {
	Zap *zap.SugaredLogger
}

// NewLogger create logger object for *gorm.DB from *echo.Logger
func NewLogger(cfg *config.EnvConfig) *Logger {

	var zapLogger *zap.Logger
	zapLogger, err := build(cfg)
	if err != nil {
		fmt.Printf("Failed to compose zap logger : %s", err)
		os.Exit(2)
	}
	sugar := zapLogger.Sugar()
	// set package variable logger.
	logger := &Logger{Zap: sugar}
	_ = zapLogger.Sync()
	return logger
}

// GetZapLogger returns zapSugaredLogger
func (log *Logger) GetZapLogger() *zap.SugaredLogger {
	return log.Zap
}

func build(cfg *config.EnvConfig) (*zap.Logger, error) {
	var zapCfg = cfg.ZapConfig
	enc, _ := newEncoder(zapCfg)
	writer, errWriter := openWriters(cfg)

	if zapCfg.Level == (zap.AtomicLevel{}) {
		return nil, errors.New("missing Level")
	}

	log := zap.New(zapcore.NewCore(enc, writer, zapCfg.Level), buildOptions(zapCfg, errWriter)...)
	return log, nil
}

func newEncoder(cfg zap.Config) (zapcore.Encoder, error) {
	switch cfg.Encoding {
	case "console":
		return zapcore.NewConsoleEncoder(cfg.EncoderConfig), nil
	case "json":
		return zapcore.NewJSONEncoder(cfg.EncoderConfig), nil
	}
	return nil, errors.New("failed to set encoder")
}

func openWriters(cfg *config.EnvConfig) (zapcore.WriteSyncer, zapcore.WriteSyncer) {
	writer := open(cfg.ZapConfig.OutputPaths, &cfg.LogRotate)
	errWriter := open(cfg.ZapConfig.ErrorOutputPaths, &cfg.LogRotate)
	return writer, errWriter
}

func open(paths []string, rotateCfg *lumberjack.Logger) zapcore.WriteSyncer {
	writers := make([]zapcore.WriteSyncer, 0, len(paths))
	for _, path := range paths {
		writer := newWriter(path, rotateCfg)
		writers = append(writers, writer)
	}
	writer := zap.CombineWriteSyncers(writers...)
	return writer
}

func newWriter(path string, rotateCfg *lumberjack.Logger) zapcore.WriteSyncer {
	switch path {
	case "stdout":
		return os.Stdout
	case "stderr":
		return os.Stderr
	}
	sink := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   path,
			MaxSize:    rotateCfg.MaxSize,
			MaxBackups: rotateCfg.MaxBackups,
			MaxAge:     rotateCfg.MaxAge,
			Compress:   rotateCfg.Compress,
		},
	)
	return sink
}

func buildOptions(cfg zap.Config, errWriter zapcore.WriteSyncer) []zap.Option {
	opts := []zap.Option{zap.ErrorOutput(errWriter)}
	if cfg.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zap.ErrorLevel
	if cfg.Development {
		stackLevel = zap.WarnLevel
	}
	if !cfg.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}
	return opts
}
