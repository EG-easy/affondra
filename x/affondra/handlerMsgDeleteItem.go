package affondra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/affondra/x/affondra/keeper"
	"github.com/EG-easy/affondra/x/affondra/types"
)

// Handle a message to delete name
func handleMsgDeleteItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteItem) (*sdk.Result, error) {
	if !k.ItemExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetItemOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteItem(ctx, msg.ID)
	return &sdk.Result{}, nil
}
