/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"code.cerinuts.io/uni/hypercocoagateway/api"
	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/spf13/cobra"
)

// gatewayCmd runs the cli in gateway mode
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Run CLI in gateway mode",
	Long: ` Run the CLI as a restful http gateway.
	This can be used e.g. to proxy between the ledger and a web application.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ac := new(api.ApiController)
		ac.Route()
		ac.Start(shared.Hyperconfig.GatewayHost, shared.Hyperconfig.GatewayPort, "")
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
}
