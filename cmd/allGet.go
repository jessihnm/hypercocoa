/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all assets on the chain",
	Long:  `Returns a list of all assets currently on the chain.`,
	Run: func(cmd *cobra.Command, args []string) {
		chain := shared.CreateChainConnection()
		assets, err := chain.GetAllAssets()
		chain.CloseConnection()
		fmt.Printf("*** Result:\n")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for _, r := range assets {
				var o []byte
				if ok, _ := cmd.Flags().GetBool("raw"); ok {
					o, _ = r.ToPrettyJson()
				} else {
					switch r.AssetType {
					case shared.AssetTypeCocoaBag:
						c := r.ToCocoaBag()
						o, _ = c.ToPrettyJson()
					case shared.AssetTypeFarm:
						c := r.ToFarm()
						o, _ = c.ToPrettyJson()
					case shared.AssetTypeFarmer:
						c := r.ToFarmer()
						o, _ = c.ToPrettyJson()
					}
				}
				fmt.Println(string(o))
			}
		}
	},
}

func init() {
	getCmd.AddCommand(allCmd)
}
