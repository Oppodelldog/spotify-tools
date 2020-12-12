package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Std LogWriter = New("Std", "Std", "Std")

func New(uuid string, name string, version string) *ZerologLogger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	return &ZerologLogger{
		l:       zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: true, TimeFormat: time.RFC3339Nano}),
		uuid:    uuid,
		name:    name,
		version: version,
	}
}

func (z ZerologLogger) withDefaultFields(e *zerolog.Event) *zerolog.Event {
	e.Str("uuid", z.uuid)
	e.Str("name", z.name)
	e.Str("version", z.version)

	return e
}

type ZerologLogger struct {
	l       zerolog.Logger
	uuid    string
	name    string
	version string
}

func (z ZerologLogger) Debugf(format string, args ...interface{}) {
	z.withDefaultFields(z.l.Debug()).Msgf(format, args...)
}

func (z ZerologLogger) Infof(format string, args ...interface{}) {
	z.withDefaultFields(z.l.Info()).Msgf(format, args...)
}

func (z ZerologLogger) Printf(format string, args ...interface{}) {
	z.l.Printf(format, args...)
}

func (z ZerologLogger) Warnf(format string, args ...interface{}) {
	z.withDefaultFields(z.l.Warn()).Msgf(format, args...)
}

func (z ZerologLogger) Warningf(format string, args ...interface{}) {
	z.withDefaultFields(z.l.Warn()).Msgf(format, args...)
}

func (z ZerologLogger) Errorf(format string, args ...interface{}) {
	z.withDefaultFields(z.l.Error()).Msgf(format, args...)
}

func (z ZerologLogger) Fatalf(format string, args ...interface{}) {
	z.withDefaultFields(z.l.Fatal()).Msgf(format, args...)
}

func (z ZerologLogger) Panicf(format string, args ...interface{}) {
	z.withDefaultFields(z.l.Panic()).Msgf(format, args...)
}

func (z ZerologLogger) Debug(msg string) {
	z.withDefaultFields(z.l.Debug()).Msg(msg)
}

func (z ZerologLogger) Info(msg string) {
	z.withDefaultFields(z.l.Info()).Msg(msg)
}

func (z ZerologLogger) Print(msg string) {
	z.l.Print(msg)
}

func (z ZerologLogger) Warn(msg string) {
	z.withDefaultFields(z.l.Warn()).Msg(msg)
}

func (z ZerologLogger) Warning(msg string) {
	z.withDefaultFields(z.l.Warn()).Msg(msg)
}

func (z ZerologLogger) Error(msg string) {
	z.withDefaultFields(z.l.Error()).Msg(msg)
}

func (z ZerologLogger) Fatal(msg string) {
	z.withDefaultFields(z.l.Fatal()).Msg(msg)
}

func (z ZerologLogger) Panic(msg string) {
	z.withDefaultFields(z.l.Panic()).Msg(msg)
}

func (z ZerologLogger) Close() {
}
