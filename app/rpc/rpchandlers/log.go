package rpchandlers

import (
	"github.com/coinexcom/kaspad/infrastructure/logger"
	"github.com/coinexcom/kaspad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
