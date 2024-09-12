package prefixmanager

import (
	"github.com/coinexcom/kaspad/infrastructure/logger"
	"github.com/coinexcom/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
