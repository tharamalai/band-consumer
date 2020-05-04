package rest

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

const (
	CdpAddrTag = "cdpAddr"
)

// RegisterRoutes - register routes for rest endpoint
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeKey string) {
	r.HandleFunc(fmt.Sprintf("/%s/cdp/{%s}", storeKey, CdpAddrTag), getCDPHandler(cliCtx, storeKey)).Methods("GET")
}
