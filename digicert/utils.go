package digicert

import (
	"go.uber.org/zap"
)

func Logger() *zap.SugaredLogger {
	return zap.S()
}
