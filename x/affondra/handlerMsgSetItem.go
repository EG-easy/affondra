package affondra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/affondra/x/affondra/keeper"
	"github.com/EG-easy/affondra/x/affondra/types"
)

func handleMsgSetItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetItem) (*sdk.Result, error) {
	var item = types.Item{
		Creator:     msg.Creator,
		ID:          msg.ID,
		Denom:       msg.Denom,
		NftId:       msg.NftId,
		Price:       msg.Price,
		Affiliate:   msg.Affiliate,
		Description: msg.Description,
		InSale:      msg.InSale,
	}
	if !msg.Creator.Equals(k.GetItemOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetItem(ctx, item)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetItem,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
	})

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
