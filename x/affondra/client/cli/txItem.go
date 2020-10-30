package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/EG-easy/affondra/x/affondra/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

//GetCmdCreateItem creates new item
func GetCmdCreateItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-item [denom] [nftId] [price] [affiliate] [description] [inSale]",
		Short: "Creates a new item",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDenom := string(args[0])
			argsNftId := string(args[1])
			argsPrice, _ := sdk.ParseCoin(args[2])
			argsAffiliate, _ := sdk.ParseCoin(args[3])
			argsDescription := string(args[4])
			argsInSale, _ := strconv.ParseBool(args[5])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateItem(cliCtx.GetFromAddress(), string(argsDenom), string(argsNftId), argsPrice, argsAffiliate, string(argsDescription), bool(argsInSale))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdSetItem sets item info
func GetCmdSetItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-item [id]  [denom] [nftId] [price] [affiliate] [description] [inSale]",
		Short: "Set a new item",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsDenom := string(args[1])
			argsNftId := string(args[2])
			argsPrice, _ := sdk.ParseCoin(args[3])
			argsAffiliate, _ := sdk.ParseCoin(args[4])
			argsDescription := string(args[5])
			argsInSale, _ := strconv.ParseBool(args[6])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetItem(cliCtx.GetFromAddress(), id, string(argsDenom), string(argsNftId), argsPrice, argsAffiliate, string(argsDescription), bool(argsInSale))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdDeleteItem deletes item
func GetCmdDeleteItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-item [id]",
		Short: "Delete a new item by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteItem(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

// GetCmdBuyItem buys item
func GetCmdBuyItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "buy-item [id] [introduced_by]",
		Short: "Buy a item by ID",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			introducedBy, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyItem(args[0], cliCtx.GetFromAddress(), introducedBy)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
