package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// MsgBuyItem degines buy item message
type MsgBuyItem struct {
	ID           string         `json:"id" yaml:"id"`
	Receiver     sdk.AccAddress `json:"receiver" yaml:"receiver"`
	IntroducedBy sdk.AccAddress `json:"introduced_by" yaml:"introduced_by"`
}

// NewMsgBuyItem is a constructor function for MsgBuyItem
func NewMsgBuyItem(id string, receiver sdk.AccAddress, introducedBy sdk.AccAddress) MsgBuyItem {
	return MsgBuyItem{
		ID:           id,
		Receiver:     receiver,
		IntroducedBy: introducedBy,
	}
}

// Route Implements Msg
func (msg MsgBuyItem) Route() string {
	return RouterKey
}

// Type Implements Msg
func (msg MsgBuyItem) Type() string {
	return "BuyItem"
}

// GetSigners Implements Msg
func (msg MsgBuyItem) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Receiver)}
}

// GetSignBytes Implements Msg
func (msg MsgBuyItem) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Implements Msg
func (msg MsgBuyItem) ValidateBasic() error {
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
