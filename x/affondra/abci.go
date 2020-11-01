package affondra

import (
	"fmt"

	"github.com/EG-easy/affondra/x/affondra/keeper"
	"github.com/EG-easy/affondra/x/affondra/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker check for infraction evidence or downtime of validators
// on every begin block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	// 	TODO: fill out if your application requires beginblock, if not you can delete this function

	k.IterateItems(ctx, func(item types.Item) (stop bool) {

		nft, _ := k.NFTKeeper.GetNFT(ctx, item.Denom, item.NftId)

		// nft owner is changed
		if !item.Creator.Equals(nft.GetOwner()) && item.InSale {

			item.ChangeInSaleStatus()
			k.SetItem(ctx, item)

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeChangeInSaleStatus,
					sdk.NewAttribute(types.AttributeKeyID, item.ID),
					sdk.NewAttribute(types.AttributeKeyReceiver, fmt.Sprintf("%t", item.InSale)),
				),
			)
		}

		return false
	})
}

// EndBlocker called every block, process inflation, update validator set.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	// 	TODO: fill out if your application requires endblock, if not you can delete this function
}
