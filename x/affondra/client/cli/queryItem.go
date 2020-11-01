package cli

import (
	"fmt"

	"github.com/EG-easy/affondra/x/affondra/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

// GetCmdListItem lists all item
func GetCmdListItem(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-item",
		Short: "list all item",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListItem, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Item\n%s\n", err.Error())
				return nil
			}
			var out []types.Item
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetItem returns specified item
func GetCmdGetItem(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "id [key]",
		Short: "Query a item by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetItem, key), nil)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Item
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetItemByOwner returns items using owner address
func GetCmdGetItemByOwner(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "addr [address]",
		Short: "Query a item by owner",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryOwner, key), nil)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Owner
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

// GetCmdGetItemByDenom returns items using denom
func GetCmdGetItemByDenom(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "denom [denom]",
		Short: "Query items by denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryDenom, key), nil)
			if err != nil {
				fmt.Printf("could not resolve item %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Collection
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
