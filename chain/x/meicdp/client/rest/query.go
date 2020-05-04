package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/tharamalai/meichain/x/meicdp/types"
)

func getCDPHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		CdpAddr := vars[CdpAddrTag]

		var cdp types.CDP

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/cdp/%s", storeName, CdpAddr),
			nil,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = cliCtx.Codec.UnmarshalJSON(res, &cdp)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, cdp)
	}
}
