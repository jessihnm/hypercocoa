package commands

import "code.cerinuts.io/uni/hypercocoagateway/shared"

// GetAsset returns a single Asset
func GetAsset(id string) (string, error) {
	chain := shared.CreateChainConnection()
	a, err := chain.ReadAssetByID(id)
	chain.CloseConnection()
	if err != nil {
		return "", err
	}
	j, err := a.ToPrettyJson()
	return string(j), err
}
