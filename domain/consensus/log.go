package consensus

import (
	"github.com/coinexcom/kaspad/infrastructure/logger"
	"github.com/coinexcom/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
