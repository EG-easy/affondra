package affondra

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/EG-easy/affondra/x/affondra/keeper"
	"github.com/EG-easy/affondra/x/affondra/types"
)

func handleMsgCreateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateItem) (*sdk.Result, error) {
	var item = types.Item{
		Creator:   msg.Creator,
		ID:        msg.ID,
		Denom:     msg.Denom,
		NftId:     msg.NftId,
		Price:     msg.Price,
		Affiliate: msg.Affiliate,
		InSale:    msg.InSale,
	}

	nft, err := k.NFTKeeper.GetNFT(ctx, msg.Denom, msg.NftId)

	//check if nft is exist
	if err != nil {
		return nil, err
	}
	// check if owner is correct
	if !msg.Creator.Equals(nft.GetOwner()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	k.CreateItem(ctx, item)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
