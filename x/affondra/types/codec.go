package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding # 1
	cdc.RegisterConcrete(MsgCreateItem{}, "affondra/CreateItem", nil)
	cdc.RegisterConcrete(MsgSetItem{}, "affondra/SetItem", nil)
	cdc.RegisterConcrete(MsgDeleteItem{}, "affondra/DeleteItem", nil)
	cdc.RegisterConcrete(MsgCreateVote{}, "affondra/CreateVote", nil)
	cdc.RegisterConcrete(MsgSetVote{}, "affondra/SetVote", nil)
	cdc.RegisterConcrete(MsgDeleteVote{}, "affondra/DeleteVote", nil)
	cdc.RegisterConcrete(MsgCreatePoll{}, "affondra/CreatePoll", nil)
	cdc.RegisterConcrete(MsgSetPoll{}, "affondra/SetPoll", nil)
	cdc.RegisterConcrete(MsgDeletePoll{}, "affondra/DeletePoll", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
