package digicert

import (
	"testing"

	"go.uber.org/zap"
)

func TestLoggerReturnsGlobalSugar(t *testing.T) {
	if logger() != zap.S() {
		t.Error("Expectd Logger() to return the global sugared logger, but returned something else")
	}
}
