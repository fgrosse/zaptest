<h1 align="center">Zaptest ⚡👩‍🔧</h1>
<p align="center">Test helpers to use a github.com/uber-go/zap logger in unit tests.</p>
<p align="center">
   <a href="https://github.com/fgrosse/zaptest/releases"><img src="https://img.shields.io/github/tag/fgrosse/zaptest.svg?label=version&color=brightgreen"></a>
   <a href="https://github.com/fgrosse/zaptest/actions/workflows/test.yml"><img src="https://github.com/fgrosse/zaptest/actions/workflows/test.yml/badge.svg"></a>
   <a href="https://goreportcard.com/report/github.com/fgrosse/zaptest"><img src="https://goreportcard.com/badge/github.com/fgrosse/zaptest"></a>
   <a href="https://pkg.go.dev/github.com/fgrosse/zaptest"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?color=blue"></a>
   <a href="https://github.com/fgrosse/zaptest/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-4183c4.svg"></a>
</p>

---

Package `zaptest` implements test helpers that facilitate using a `zap.Logger`
in standard go unit tests.

This package is useful when running unit tests with components that use the
https://github.com/uber-go/zap logging framework. In unit tests we usually want
to suppress any logging output as long as the tests are not failing or they are
started in a verbose mode.

## Installation 

```sh
$ go get github.com/fgrosse/zaptest
```

## Usage with standard unit tests

You can use `zaptest` in standard unit tests. All log output will be sent via
the `testing.T` so it will only be shown if the test fails or if the `-v` flag
was set.

Note that you can also use `zaptest` in benchmarks because `testing.B` also
implements the necessary logger interface.

```go
func TestLogger(t *testing.T) {
	l := zaptest.Logger(t)
	l.Debug("This is a debug message, debug messages will be logged as well")
	l.Info("Logs will not be shown during normal test execution")
	l.Warn("You can see all log messages of a successful run by passing the -v flag")
	l.Error("Additionally the entire log output for a specific unit test will be visible when a test fails")
}
```

## Usage with ginkgo

Package `zaptest` is also compatible with the https://github.com/onsi/ginkgo BDD
testing framework. As with the standard unit tests, any log output for a
specific test will only be printed if that test fails or if the test is running
in verbose mode. With ginkgo use `ginkgo -v` to enable verbose output (`go test -v`
does not seem to work).

```go
// TestedType is an example of some type you want to test.
// Note that the TestedType uses a zap logger.
type TestedType struct {
	log *zap.Logger
}

// DoStuff is an example function of the TestedType which uses a logger.
func (tt *TestedType) DoStuff() error {
	tt.log.Debug("Doing stuff")
	return nil
}

// TestLoggerWriter shows how to use the LoggerWriter function for ginkgo tests.
// Run `ginkgo -v`to see the output of this test.
func TestLoggerWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Example Suite")
}

// Describe the TestedType using ginkgo/gomega.
var _ = Describe("TestedType", func() {
	It("should do stuff", func() {
		tt := &TestedType{log: zaptest.LoggerWriter(GinkgoWriter)}
		Expect(tt.DoStuff()).To(Succeed())
	})
})
```

## Usage as library

You can also use `zaptest` as library since it does not import the `testing`
package.
