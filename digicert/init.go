package digicert

import (
	"go.uber.org/zap"
)

func init() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}
