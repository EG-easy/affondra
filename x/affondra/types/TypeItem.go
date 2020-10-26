package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Item struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ID        string         `json:"id" yaml:"id"`
	Denom     string         `json:"denom" yaml:"denom"`
	NftId     string         `json:"nftId" yaml:"nftId"`
	Price     sdk.Coin       `json:"price" yaml:"price"`
	Affiliate sdk.Coin       `json:"affiliate" yaml:"affiliate"`
	InSale    bool           `json:"inSale" yaml:"inSale"`
}

func NewItem(id string, owner sdk.AccAddress, denom string, nftId string, price sdk.Coin, affiliate sdk.Coin, inSale bool) Item {
	return Item{
		Creator:   owner,
		ID:        id,
		Denom:     denom,
		NftId:     nftId,
		Price:     price,
		Affiliate: affiliate,
		InSale:    inSale,
	}
}

func (item Item) GetID() string {
	return item.ID
}

func (item Item) GetOwner() sdk.AccAddress {
	return item.Creator
}

func (item Item) SetOwner(addr sdk.AccAddress) {
	item.Creator = addr
}

func (item Item) String() string {
	return fmt.Sprintf(`ID:%s
Creator:%s
Denom:%s
NftId:%s
Price:%s
Affiliate:%s
InSale:%v`,
		item.ID,
		item.Creator,
		item.Denom,
		item.NftId,
		item.Price,
		item.Affiliate,
		item.InSale,
	)
}

//func (item Item) find(id string) int {
//	return FindUtil(item, id)
//}
