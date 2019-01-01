// Package zaptest implements test helpers that facilitate using a zap.Logger in
// unit tests.
//
// This package is useful when running unit tests with components that use the
// github.com/uber-go/zap logging framework. In unit tests we usually want to
// suppress any logging output as long as the tests are not failing or they are
// started in a verbose mode.
//
// Package zaptest is also compatible with the github.com/onsi/ginkgo BDD
// testing framework. Have a look at the zaptest unit tests to see an usage
// examples.
package zaptest

import (
	"io"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger is the minimal subset of logging functions that exists both in
// testing.T and testing.B.
type logger interface {
	Log(args ...interface{})
}

// testOutput is a zapcore.SyncWriter that writes all output to l.
type testOutput struct {
	logger
}

// writeSyncer decorates an io.Writer with a no-op Sync() function.
type writeSyncer struct {
	io.Writer
}

// Config is the default function used create the configuration for the zap
// logger. You can change this if you want to change the logger encoding options.
//
// By default the zap development config will be used and timestamps will be
// omitted (just like normal t.Log does).
var Config = func() zapcore.EncoderConfig {
	conf := zap.NewDevelopmentEncoderConfig()

	// In unit tests we are not that interested in full timestamps so we omit
	// them just like testing.T.Log does.
	conf.TimeKey = ""

	return conf
}

// Level is the zap log level that will be printed in unit tests.
var Level = zap.DebugLevel

// Logger creates a new zap.Logger that writes all messages via t.Log(…).
// Note that both testing.T and testing.B implement the logger interface.
func Logger(t logger) *zap.Logger {
	return newLogger(writeSyncer{testOutput{t}})
}

// LoggerWriter creates a new zap.Logger that writes all messages to the given
// io.Writer.
func LoggerWriter(w io.Writer) *zap.Logger {
	return newLogger(writeSyncer{w})
}

// newLogger creates a *new zap.Logger using the package level Config function
// and Level value.
func newLogger(w zapcore.WriteSyncer) *zap.Logger {
	conf := Config()
	enc := zapcore.NewConsoleEncoder(conf)
	core := zapcore.NewCore(enc, w, Level)

	return zap.New(core)
}

// Write logs all messages as logs via o.
func (o testOutput) Write(p []byte) (int, error) {
	msg := strings.TrimSpace(string(p))
	o.Log(msg)
	return len(p), nil
}

// Sync does nothing since all output was written to the writer immediately.
func (ws writeSyncer) Sync() error {
	return nil
}
