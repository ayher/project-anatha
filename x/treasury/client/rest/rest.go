package rest

import (
	"github.com/gorilla/mux"

	"github.com/ayher/anatha/client/context"
)

// RegisterRoutes registers treasury-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	// registerTxRoutes(cliCtx, r)
}
