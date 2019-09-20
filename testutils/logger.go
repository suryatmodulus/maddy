package testutils

import (
	"strings"
	"testing"
	"time"

	"github.com/foxcpp/maddy/log"
)

func Logger(t *testing.T, name string) log.Logger {
	if testing.Verbose() {
		return log.Logger{
			Out: log.FuncOutput(func(_ time.Time, debug bool, str string) {
				t.Helper()
				str = strings.TrimSuffix(str, "\n")
				if debug {
					str = "[debug] " + str
				}
				t.Log(str)
			}, func() error {
				return nil
			}),
			Name:  name,
			Debug: true,
		}
	}

	// MultiOutput to empty slice is a blackhole.
	return log.Logger{Out: log.MultiOutput()}
}
