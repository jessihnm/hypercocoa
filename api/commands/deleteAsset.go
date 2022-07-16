package commands

import "code.cerinuts.io/uni/hypercocoagateway/shared"

// DeleteAsset removes an asset from the chain
func DeleteAsset(id string) (string, error) {
	chain := shared.CreateChainConnection()
	a, err := chain.DeleteAsset(id)
	chain.CloseConnection()
	return a, err
}
