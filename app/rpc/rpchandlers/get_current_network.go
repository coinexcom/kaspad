package rpchandlers

import (
	"github.com/coinexcom/kaspad/app/appmessage"
	"github.com/coinexcom/kaspad/app/rpc/rpccontext"
	"github.com/coinexcom/kaspad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
