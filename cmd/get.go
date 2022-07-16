/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get things from the ledger",
	Long: `Get any asset from the ledger. 
	Currently only assets and metadata are supported`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if id == "" || err != nil {
			fmt.Println("No Asset ID provided")
		}
		chain := shared.CreateChainConnection()
		a, err := chain.ReadAssetByID(id)
		chain.CloseConnection()
		if err != nil {
			return
		}
		result, _ := a.ToPrettyJson()

		fmt.Printf("*** Result:%s\n", result)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().String("id", "", "The id of the asset to remove")
	getCmd.PersistentFlags().Bool("raw", false, "Print raw assets")
}
