package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers affondra-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/affondra/item", createItemHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/affondra/item", listItemHandler(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/item/owner/{ownerAddr}", getItemByOwner(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/item/denom/{denom}", getItemByDenom(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/item/{id}", getItemById(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/item", setItemHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/affondra/item", deleteItemHandler(cliCtx)).Methods("DELETE")
}
