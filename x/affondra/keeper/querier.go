package keeper

import (
	// this line is used by starport scaffolding # 1
	"fmt"

	"github.com/EG-easy/affondra/x/affondra/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for affondra clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		fmt.Printf("path from New Querier: %s", path)
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryOwner:
			return queryOwner(ctx, path[1:], req, k)
		case types.QueryListItem:
			return listItem(ctx, k)
		case types.QueryGetItem:
			return getItem(ctx, path[1:], k)
		case types.QueryListVote:
			return listVote(ctx, k)
		case types.QueryGetVote:
			return getVote(ctx, path[1:], k)
		case types.QueryListPoll:
			return listPoll(ctx, k)
		case types.QueryGetPoll:
			return getPoll(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown affondra query endpoint")
		}
	}
}

func queryOwner(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryOwnerParams

	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, err.Error())
	}

	owner := k.GetOwner(ctx, params.Owner)
	bz, err := types.ModuleCdc.MarshalJSON(owner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
