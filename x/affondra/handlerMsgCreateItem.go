package affondra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/affondra/x/affondra/keeper"
	"github.com/EG-easy/affondra/x/affondra/types"
)

func handleMsgCreateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateItem) (*sdk.Result, error) {

	item := types.NewItem(msg.ID, msg.Creator, msg.Denom, msg.NftId, msg.Price, msg.Affiliate, msg.InSale)

	nft, err := k.NFTKeeper.GetNFT(ctx, msg.Denom, msg.NftId)

	//check if nft is exist
	if err != nil {
		return nil, err
	}
	// check if owner is correct
	if !msg.Creator.Equals(nft.GetOwner()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	//check if affiliate <= price
	if msg.Price.IsLT(msg.Affiliate) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAffiliatePrice, "Incorrect affiliate price")
	}
	//check if on sale
	//	if msg.InSale {
	//		return nil, sdkerrors.Wrap(types.ErrAlreadyOnSale, "Alrady on sale")
	//	}

	k.CreateItem(ctx, item)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateItem,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
