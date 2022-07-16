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

// farmerCmd represents the farm command
var farmerCmd = &cobra.Command{
	Use:   "farmer",
	Short: "Add a farmer",
	Long: `Add a farmer to the chain.
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

		firstname, err := cmd.Flags().GetString("firstname")
		cobra.CheckErr(err)
		if firstname == "" {
			os.Stderr.WriteString("No firstname specified")
			return
		}

		surname, err := cmd.Flags().GetString("surname")
		cobra.CheckErr(err)
		if surname == "" {
			os.Stderr.WriteString("No surname specified")
			return
		}

		email, err := cmd.Flags().GetString("email")
		cobra.CheckErr(err)
		if email == "" {
			os.Stderr.WriteString("No email specified")
			return
		}

		phone, err := cmd.Flags().GetString("phone")
		cobra.CheckErr(err)
		if phone == "" {
			os.Stderr.WriteString("No phone specified")
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

		farmer := shared.Farmer{
			ID: id,
			Address: shared.Address{
				Number:  number,
				Street:  street,
				City:    city,
				ZipCode: zipcode,
				Country: country,
			},
			Surname:   surname,
			Firstname: firstname,
			Phone:     phone,
			Email:     email,
		}
		var j []byte
		if human {
			j, err = farmer.ToPrettyJson()

		} else {
			j, err = farmer.ToJson()
		}
		cobra.CheckErr(err)
		fmt.Println(string(j))

		chain := shared.CreateChainConnection()
		chain.CreateAsset(farmer.ToAsset())
		chain.CloseConnection()
	},
}

func init() {
	addCmd.AddCommand(farmerCmd)

	farmerCmd.Flags().StringP("number", "u", "", "The house number of the farmer's house")
	farmerCmd.Flags().StringP("street", "s", "", "The street name of the farmer's house")
	farmerCmd.Flags().StringP("city", "y", "", "The city of the farmer's house")
	farmerCmd.Flags().StringP("zipcode", "z", "", "The zip code of the farmer's house")
	farmerCmd.Flags().StringP("country", "c", "", "The country of the farmer's house")
	farmerCmd.Flags().StringP("firstname", "f", "", "The first name of the farmer")
	farmerCmd.Flags().StringP("surname", "n", "", "The surname of the farmer")
	farmerCmd.Flags().StringP("phone", "p", "", "The phone of the farmer")
	farmerCmd.Flags().StringP("email", "e", "", "The email of the farmer")
}
