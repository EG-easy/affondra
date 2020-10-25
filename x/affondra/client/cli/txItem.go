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

func GetCmdCreateItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-item [denom] [nftId] [price] [affiliate] [inSale]",
		Short: "Creates a new item",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDenom := string(args[0])
			argsNftId := string(args[1])
			argsPrice, _ := sdk.ParseCoin(args[2])
			argsAffiliate, _ := sdk.ParseCoin(args[3])
			argsInSale, _ := strconv.ParseBool(args[4])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateItem(cliCtx.GetFromAddress(), string(argsDenom), string(argsNftId), argsPrice, argsAffiliate, bool(argsInSale))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdSetItem(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-item [id]  [denom] [nftId] [price] [affiliate] [inSale]",
		Short: "Set a new item",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsDenom := string(args[1])
			argsNftId := string(args[2])
			argsPrice, _ := sdk.ParseCoin(args[3])
			argsAffiliate, _ := sdk.ParseCoin(args[4])
			argsInSale, _ := strconv.ParseBool(args[5])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetItem(cliCtx.GetFromAddress(), id, string(argsDenom), string(argsNftId), argsPrice, argsAffiliate, bool(argsInSale))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

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
