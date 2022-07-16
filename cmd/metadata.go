/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
	"github.com/spf13/cobra"
)

// metadataCmd prints the metadata
var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Print metadata of Chaincode",
	Long:  `Print the metadata of the chaincode by invoking org.hyperledger.fabric:GetMetadata`,
	Run: func(cmd *cobra.Command, args []string) {
		chain := shared.CreateChainConnection()
		result, _ := chain.GetMetadata()
		chain.CloseConnection()
		fmt.Printf("*** Result:%s\n", result)
	},
}

func init() {
	getCmd.AddCommand(metadataCmd)
}
