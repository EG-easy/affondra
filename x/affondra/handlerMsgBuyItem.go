package affondra

import (
	"github.com/EG-easy/affondra/x/affondra/keeper"
	"github.com/EG-easy/affondra/x/affondra/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func handleMsgBuyItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgBuyItem) (*sdk.Result, error) {

	item, err := k.GetItem(ctx, msg.ID)
	// check if item is exsit
	if err != nil {
		return nil, err
	}

	nft, err := k.NFTKeeper.GetNFT(ctx, item.Denom, item.NftId)

	//check if nft is exist
	if err != nil {
		return nil, err
	}

	// check if owner is correct
	if !item.Creator.Equals(nft.GetOwner()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	//check if on sale
	if !item.InSale {
		return nil, sdkerrors.Wrap(types.ErrOutOfSale, "Alrady out of sale")
	}

	// check if the receiver has enough amount
	coins := k.CoinKeeper.GetCoins(ctx, msg.Receiver)
	if !coins.AmountOf(item.Price.Denom).GT(item.Price.Amount) {
		return nil, sdkerrors.Wrap(types.ErrInvalidAffiliatePrice, "Incorrect affiliate price")
	}

	// update item receiver/insale info
	item.SetReceiver(msg.Receiver)
	item.ChangeInSaleStatus()
	k.SetItem(ctx, item)

	// update NFT owner
	nft.SetOwner(msg.Receiver)
	// update the NFT (owners are updated within the keeper)
	err = k.NFTKeeper.UpdateNFT(ctx, item.Denom, nft)
	if err != nil {
		return nil, err
	}
	// payment to creator
	payment := sdk.NewCoins(item.Price)
	if err := k.CoinKeeper.SendCoins(ctx, msg.Receiver, item.Creator, payment); err != nil {
		return nil, err
	}
	// pay for referral fee
	referralFee := sdk.NewCoins(item.Price)
	if err := k.CoinKeeper.SendCoins(ctx, msg.Receiver, msg.IntroducedBy, referralFee); err != nil {
		return nil, err
	}

	//:TODO adding event here

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
