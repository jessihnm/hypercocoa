package commands

import (
	"encoding/json"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
)

// GetAllAssets returns a list of all Assets
func GetAllAssets() (string, error) {
	chain := shared.CreateChainConnection()
	a, err := chain.GetAllAssets()
	chain.CloseConnection()
	if err != nil {
		return "", err
	}
	j, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "", err
	}
	return string(j), err
}
