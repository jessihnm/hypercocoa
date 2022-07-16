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

// farmCmd represents the farm command
var farmCmd = &cobra.Command{
	Use:   "farm",
	Short: "Add a farm",
	Long: `Add a farm to the chain.
	The Farmer reference should be known.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		human, _ := cmd.Flags().GetBool("human")

		if human {
			fmt.Println("Adding a farm...")
		}

		id, err := cmd.Flags().GetString("id")
		cobra.CheckErr(err)
		if id == "" {
			id = uuid.NewString()
		}

		size, err := cmd.Flags().GetInt("size")
		cobra.CheckErr(err)
		if size == -1 {
			os.Stderr.WriteString("No size specified")
			return
		}

		farmer, err := cmd.Flags().GetString("farmer")
		cobra.CheckErr(err)
		if farmer == "" {
			os.Stderr.WriteString("No farmer specified")
			return
		}
		_, err = uuid.Parse(farmer)
		if err != nil {
			os.Stderr.WriteString("Invalid farmer reference specified")
			return
		}

		number, err := cmd.Flags().GetString("number")
		cobra.CheckErr(err)
		if number == "" {
			os.Stderr.WriteString("No number specified")
			return
		}

		street, err := cmd.Flags().GetString("street")
		cobra.CheckErr(err)
		if street == "" {
			os.Stderr.WriteString("No street specified")
			return
		}

		city, err := cmd.Flags().GetString("city")
		cobra.CheckErr(err)
		if city == "" {
			os.Stderr.WriteString("No city specified")
			return
		}

		zipcode, err := cmd.Flags().GetString("zipcode")
		cobra.CheckErr(err)
		if zipcode == "" {
			os.Stderr.WriteString("No zipcode specified")
			return
		}

		country, err := cmd.Flags().GetString("country")
		cobra.CheckErr(err)
		if country == "" {
			os.Stderr.WriteString("No country specified")
			return
		}

		// this might be empty
		name, _ := cmd.Flags().GetString("name")

		farm := shared.Farm{
			ID:   id,
			Size: size,
			Address: shared.Address{
				Number:  number,
				Street:  street,
				City:    city,
				ZipCode: zipcode,
				Country: country,
			},
			Name:            name,
			FarmerReference: farmer,
		}
		var j []byte
		if human {
			j, err = farm.ToPrettyJson()

		} else {
			j, err = farm.ToJson()
		}
		cobra.CheckErr(err)
		fmt.Println(string(j))

		chain := shared.CreateChainConnection()
		chain.CreateAsset(farm.ToAsset())
		chain.CloseConnection()
	},
}

func init() {
	addCmd.AddCommand(farmCmd)

	farmCmd.Flags().IntP("size", "w", -1, "The size of the farm in ha")
	farmCmd.Flags().StringP("farmer", "f", "", "The farmer reference id")
	farmCmd.Flags().StringP("number", "u", "", "The house number of the farm")
	farmCmd.Flags().StringP("street", "s", "", "The street name of the farm")
	farmCmd.Flags().StringP("city", "y", "", "The city of the farm")
	farmCmd.Flags().StringP("zipcode", "z", "", "The zip code of the farm")
	farmCmd.Flags().StringP("country", "c", "", "The country of the farm")
	farmCmd.Flags().StringP("name", "n", "", "The name of the farm")
}
