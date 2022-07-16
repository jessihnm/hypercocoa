/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/spf13/cobra"
)

// printconfigCmd prints the current config
var printconfigCmd = &cobra.Command{
	Use:   "printconfig",
	Short: "Print config",
	Long:  `Print the currently used config`,
	Run: func(cmd *cobra.Command, args []string) {
		shared.PrintConfig()
	},
}

func init() {
	rootCmd.AddCommand(printconfigCmd)
}
