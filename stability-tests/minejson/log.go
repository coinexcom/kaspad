package main

import (
	"github.com/coinexcom/kaspad/infrastructure/logger"
	"github.com/coinexcom/kaspad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("MNJS")
	spawn      = panics.GoroutineWrapperFunc(log)
)
