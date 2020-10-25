package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateItem{}

type MsgCreateItem struct {
	ID        string
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	Denom     string         `json:"denom" yaml:"denom"`
	NftId     string         `json:"nftId" yaml:"nftId"`
	Price     sdk.Coin       `json:"price" yaml:"price"`
	Affiliate int32          `json:"affiliate" yaml:"affiliate"`
	InSale    bool           `json:"inSale" yaml:"inSale"`
}

func NewMsgCreateItem(creator sdk.AccAddress, denom string, nftId string, price sdk.Coin, affiliate int32, inSale bool) MsgCreateItem {
	return MsgCreateItem{
		ID:        uuid.New().String(),
		Creator:   creator,
		Denom:     denom,
		NftId:     nftId,
		Price:     price,
		Affiliate: affiliate,
		InSale:    inSale,
	}
}

func (msg MsgCreateItem) Route() string {
	return RouterKey
}

func (msg MsgCreateItem) Type() string {
	return "CreateItem"
}

func (msg MsgCreateItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateItem) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
