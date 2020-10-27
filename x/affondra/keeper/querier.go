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
		fmt.Printf("path from New Querier: %s\n", path)
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryOwner:
			return queryOwner(ctx, path[1:], req, k)
		case types.QueryDenom:
			return queryCollection(ctx, path[1:], req, k)
		case types.QueryListItem:
			return listItem(ctx, k)
		case types.QueryGetItem:
			return getItem(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown affondra query endpoint")
		}
	}
}

func queryOwner(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	fmt.Print("===start query owner===\n")
	fmt.Printf("request: %v\n", req)
	fmt.Printf("path: %s\n", path)

	address, err := sdk.AccAddressFromBech32(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	owner, _ := k.GetOwner(ctx, address)
	bz, err := types.ModuleCdc.MarshalJSON(owner)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func queryCollection(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	fmt.Print("===start query collection===\n")
	fmt.Printf("request: %v\n", req)
	fmt.Printf("path: %s\n", path)

	denom := path[0]
	collection, found := k.GetCollection(ctx, denom)
	fmt.Printf("collection: %v\n", collection)
	fmt.Printf("found: %v\n", found)

	if !found {
		return nil, sdkerrors.Wrap(types.ErrUnknownCollection, fmt.Sprintf("unknown denom %s", denom))
	}

	bz, err := types.ModuleCdc.MarshalJSON(collection)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
