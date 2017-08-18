package zaptest_test

import (
	"testing"

	"github.com/fgrosse/zaptest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// TestLogger shows how to use a zaptest with normal unit tests.
// Run `go test -v`to see the output of this test.
func TestLogger(t *testing.T) {
	l := zaptest.Logger(t)
	l.Debug("This is a debug message, debug messages will be logged as well")
	l.Info("Logs will not be shown during normal test execution")
	l.Warn("You can see all log messages of a successful run by passing the -v flag")
	l.Error("Additionally the entire log output for a specific unit test will be visible when a test fails")
}

// BenchmarkLogger demonstrates that logging in standard benchmarks works just
// like logging in normal unit tests.
func BenchmarkLogger(b *testing.B) {
	l := zaptest.Logger(b)
	l.Info("Logging in benchmarks works the same way")
}

// TestLoggerWriter show how to use the LoggerWriter function for ginkgo tests.
// Run `ginkgo -v`to see the output of this test.
func TestLoggerWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo Example Suite")
}

// TestedType is an example of some type you want to test.
type TestedType struct {
	log *zap.Logger
}

// DoStuff is an example function of the TestedType which uses a logger.
func (tt *TestedType) DoStuff() error {
	tt.log.Debug("Doing stuff")
	return nil
}

// Describe the TestedType using ginkgo/gomega.
var _ = Describe("TestedType", func() {
	It("should do stuff", func() {
		tt := &TestedType{log: zaptest.LoggerWriter(GinkgoWriter)}
		Expect(tt.DoStuff()).To(Succeed())
	})
})
