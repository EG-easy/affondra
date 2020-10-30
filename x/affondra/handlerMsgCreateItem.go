package affondra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/affondra/x/affondra/keeper"
	"github.com/EG-easy/affondra/x/affondra/types"
)

func handleMsgCreateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateItem) (*sdk.Result, error) {

	nft, err := k.NFTKeeper.GetNFT(ctx, msg.Denom, msg.NftId)
	//check if nft is exist
	if err != nil {
		return nil, err
	}
	//check if the item already exsits
	exsits := k.ItemExists(ctx, msg.ID)
	if exsits {
		return nil, sdkerrors.Wrap(types.ErrAlreadyOnSale, "The item is already on market")
	}

	item := types.NewItem(msg.ID, msg.Creator, msg.Denom, msg.NftId, nft.GetTokenURI(), msg.Price, msg.Affiliate, msg.Description, msg.InSale)
	// check if owner is correct
	if !msg.Creator.Equals(nft.GetOwner()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	//check if affiliate <= price
	if msg.Price.IsLT(msg.Affiliate) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAffiliatePrice, "Incorrect affiliate price")
	}

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
