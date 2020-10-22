package voter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/EG-easy/voter/x/voter/types"
	"github.com/EG-easy/voter/x/voter/keeper"
)

func handleMsgCreateItem(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateItem) (*sdk.Result, error) {
	var item = types.Item{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Denom: msg.Denom,
    	NftId: msg.NftId,
    	Price: msg.Price,
    	Affiliate: msg.Affiliate,
    	InSale: msg.InSale,
	}
	k.CreateItem(ctx, item)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
