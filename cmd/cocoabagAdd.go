/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// cocoabagCmd adds a cocoabag
var cocoabagAddCmd = &cobra.Command{
	Use:   "cocoabag",
	Short: "Add a cocoabag",
	Long: `Add a bag of cocoa to the chain.
	The Farm reference should be known.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		human, _ := cmd.Flags().GetBool("human")

		if human {
			fmt.Println("Adding a fresh bag of cocoa...")
		}

		id, err := cmd.Flags().GetString("id")
		cobra.CheckErr(err)
		if id == "" {
			id = uuid.NewString()
		}

		cocoatype, err := cmd.Flags().GetString("type")
		cobra.CheckErr(err)
		if cocoatype == "" {
			os.Stderr.WriteString("No type specified")
			return
		}

		weight, err := cmd.Flags().GetInt("weight")
		cobra.CheckErr(err)
		if weight == -1 {
			os.Stderr.WriteString("No weight specified")
			return
		}

		farm, err := cmd.Flags().GetString("farm")
		cobra.CheckErr(err)
		if farm == "" {
			os.Stderr.WriteString("No farm specified")
			return
		}
		_, err = uuid.Parse(farm)
		if err != nil {
			os.Stderr.WriteString("Invalid farm reference specified")
			return
		}

		price, err := cmd.Flags().GetInt("price")
		cobra.CheckErr(err)
		if price == -1 {
			os.Stderr.WriteString("No price specified")
			return
		}

		currency, err := cmd.Flags().GetString("currency")
		cobra.CheckErr(err)
		if currency == "" {
			os.Stderr.WriteString("No currency specified")
			return
		}

		bag := shared.CocoaBag{
			ID:            id,
			Type:          cocoatype,
			Weight:        weight,
			FarmReference: farm,
			Price:         price,
			Currency:      currency,
		}
		var j []byte
		if human {
			j, err = bag.ToPrettyJson()

		} else {
			j, err = bag.ToJson()
		}
		cobra.CheckErr(err)
		fmt.Println(string(j))

		chain := shared.CreateChainConnection()
		chain.CreateAsset(bag.ToAsset())
		chain.CloseConnection()
	},
}

func init() {
	addCmd.AddCommand(cocoabagAddCmd)

	cocoabagAddCmd.Flags().IntP("weight", "w", -1, "The weight of the bag in kg")
	cocoabagAddCmd.Flags().StringP("type", "t", "", "The cocoa type")
	cocoabagAddCmd.Flags().StringP("farm", "f", "", "The farm reference id")
	cocoabagAddCmd.Flags().IntP("price", "p", -1, "The price of the bag")
	cocoabagAddCmd.Flags().StringP("currency", "c", "USD", "The currency of the price")
}
