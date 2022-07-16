/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/spf13/cobra"
)

// cocoabagUpdateCmd represents the cocoabag command
var cocoabagUpdateCmd = &cobra.Command{
	Use:   "cocoabag",
	Short: "Update a cocoabag",
	Long:  `Update cocoabag.`,
	Run: func(cmd *cobra.Command, args []string) {
		human, _ := cmd.Flags().GetBool("human")

		if human {
			fmt.Println("Updating your bag of cocoa...")
		}

		id, err := cmd.Flags().GetString("id")
		cobra.CheckErr(err)
		if id == "" {
			os.Stderr.WriteString("No id specified")
			return
		}

		weight, err := cmd.Flags().GetInt("weight")
		factory, err := cmd.Flags().GetString("factory")
		price, err := cmd.Flags().GetInt("price")
		currency, err := cmd.Flags().GetString("currency")

		chain := shared.CreateChainConnection()
		a, err := chain.ReadAssetByID(id)
		cb := a.ToCocoaBag()

		if factory != "" {
			cb.FactoryReference = factory
		}

		if weight != -1 {
			cb.Weight = weight
		}

		if price != -1 {
			cb.Price = price
		}

		if currency != "" {
			cb.Currency = currency
		}

		var j []byte
		if human {
			j, err = cb.ToPrettyJson()

		} else {
			j, err = cb.ToJson()
		}
		cobra.CheckErr(err)
		fmt.Println(string(j))

		chain.UpdateAsset(cb.ToAsset())
		chain.CloseConnection()
	},
}

func init() {
	updateCmd.AddCommand(cocoabagUpdateCmd)

	cocoabagUpdateCmd.Flags().IntP("weight", "w", -1, "The weight of the bag in kg")
	cocoabagUpdateCmd.Flags().StringP("factory", "f", "", "The farm reference id")
	cocoabagUpdateCmd.Flags().IntP("price", "p", -1, "The price of the bag")
	cocoabagUpdateCmd.Flags().StringP("currency", "c", "USD", "The currency of the price")
}
