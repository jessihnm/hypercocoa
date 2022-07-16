/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/spf13/cobra"
)

// assetRmCmd represents the asset command
var assetRmCmd = &cobra.Command{
	Use:   "asset",
	Short: "Remove an asset from the ledger",
	Long:  `Remove asset from the ledger by given ID`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if id == "" || err != nil {
			fmt.Println("No Asset ID provided")
		}
		chain := shared.CreateChainConnection()
		result, _ := chain.DeleteAsset(id)
		chain.CloseConnection()
		fmt.Printf("*** Result:%s\n", result)
	},
}

func init() {
	rmCmd.AddCommand(assetRmCmd)

	assetRmCmd.Flags().String("id", "", "The id of the asset to remove")
}
