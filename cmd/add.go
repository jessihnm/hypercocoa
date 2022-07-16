/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add things to the chain",
	Long:  `Add an asset to the chain. Currently only cocoabags are supported`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().String("id", "", "Set an explicit ID")
	addCmd.PersistentFlags().Bool("human", false, "Pretty print to be read by a human")
}
