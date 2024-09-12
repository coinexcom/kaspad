package connmanager

import (
	"github.com/coinexcom/kaspad/infrastructure/logger"
	"github.com/coinexcom/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
