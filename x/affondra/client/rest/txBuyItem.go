package rest

import (
	"net/http"

	"github.com/EG-easy/affondra/x/affondra/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

type buyItemRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	ID           string       `json:"ID"`
	Receiver     string       `json:"receiver"`
	IntroducedBy string       `json:"introducedBy"`
}

func buyItemHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req buyItemRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		receiver, err := sdk.AccAddressFromBech32(req.Receiver)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		introducedBy, err := sdk.AccAddressFromBech32(req.IntroducedBy)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgBuyItem(
			req.ID,
			receiver,
			introducedBy,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
