package logger

// thx zap  sugar
import (
	"fmt"
	"time"

	"github.com/exuan/waka-api/internal/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
)

type (
	Config struct {
		Level  zapcore.Level
		Name   string
		Format string
	}

	Sugar struct {
		base *zap.Logger
	}
)

func New(cfg *config.Config, configs ...*Config) (*Sugar, error) {
	cores := make([]zapcore.Core, 0)
	encoder := zapcore.EncoderConfig{
		TimeKey:     "time",
		NameKey:     "logger",
		CallerKey:   "caller",
		MessageKey:  "",
		LineEnding:  zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	for _, c := range configs {
		conf := c
		link := fmt.Sprintf("%s%s.%s.log", cfg.GetString("logPath"), cfg.SrvName, conf.Name)
		path := fmt.Sprintf("%s%s.%s.%s.log", cfg.GetString("logPath"), cfg.SrvName, conf.Name, conf.Format)

		rl, err := rotatelogs.New(path, rotatelogs.WithLinkName(link))
		if err != nil {
			return nil, err
		}

		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder),
			zapcore.AddSync(rl),
			zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
				return lev == conf.Level
			}),
		))
	}

	// like zl.Sugar()
	zl := zap.New(zapcore.NewTee(cores...)).WithOptions(zap.AddCallerSkip(2))
	return &Sugar{zl}, nil
}

func (s *Sugar) Write(lvl zapcore.Level, msg string, keysAndValues ...interface{}) {
	s.log(lvl, msg, nil, keysAndValues)
}

func (s *Sugar) Sync() error {
	return s.base.Sync()
}

func (s *Sugar) log(lvl zapcore.Level, template string, fmtArgs []interface{}, context []interface{}) {
	s.base.WithOptions()
	if lvl < zap.DPanicLevel && !s.base.Core().Enabled(lvl) {
		return
	}

	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	if ce := s.base.Check(lvl, msg); ce != nil {
		ce.Write(s.sweetenFields(context)...)
	}
}

func (s *Sugar) sweetenFields(args []interface{}) []zap.Field {
	if len(args) == 0 {
		return nil
	}

	// Allocate enough space for the worst case; if users pass only structured
	// fields, we shouldn't penalize them with extra allocations.
	fields := make([]zap.Field, 0, len(args))
	var invalid invalidPairs

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			s.base.DPanic(_oddNumberErrMsg, zap.Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		s.base.DPanic(_nonStringKeyErrMsg, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	zap.Any("key", p.key).AddTo(enc)
	zap.Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}
