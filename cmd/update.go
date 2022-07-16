/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an asset",
	Long:  `Update an asset, e.g. a cocoabag`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().String("id", "", "The id of the asset to remove")
}
