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
	r.HandleFunc("/affondra/item/{key}", getItemHandler(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/item", setItemHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/affondra/item", deleteItemHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc("/affondra/vote", createVoteHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/affondra/vote", listVoteHandler(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/vote/{key}", getVoteHandler(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/vote", setVoteHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/affondra/vote", deleteVoteHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc("/affondra/poll", createPollHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/affondra/poll", listPollHandler(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/poll/{key}", getPollHandler(cliCtx, "affondra")).Methods("GET")
	r.HandleFunc("/affondra/poll", setPollHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/affondra/poll", deletePollHandler(cliCtx)).Methods("DELETE")

}
