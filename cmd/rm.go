/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rmCmd is used to remove things from the ledger
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an asset, requires ID",
	Long:  `Remove things from the chain, currently only assets are supported`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
