package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetItem{}

type MsgSetItem struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Denom string `json:"denom" yaml:"denom"`
  NftId string `json:"nftId" yaml:"nftId"`
  Price int32 `json:"price" yaml:"price"`
  Affiliate int32 `json:"affiliate" yaml:"affiliate"`
  InSale bool `json:"inSale" yaml:"inSale"`
}

func NewMsgSetItem(creator sdk.AccAddress, id string, denom string, nftId string, price int32, affiliate int32, inSale bool) MsgSetItem {
  return MsgSetItem{
    ID: id,
		Creator: creator,
    Denom: denom,
    NftId: nftId,
    Price: price,
    Affiliate: affiliate,
    InSale: inSale,
	}
}

func (msg MsgSetItem) Route() string {
  return RouterKey
}

func (msg MsgSetItem) Type() string {
  return "SetItem"
}

func (msg MsgSetItem) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetItem) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetItem) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}