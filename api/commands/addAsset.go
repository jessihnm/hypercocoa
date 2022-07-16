package commands

import (
	"encoding/json"
	"fmt"
	"io"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
)

// AddAsset adds an asset to the chain
func AddAsset(input io.ReadCloser) (string, error) {
	a := shared.Asset{}
	err := json.NewDecoder(input).Decode(&a)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	x, _ := a.ToPrettyJson()
	fmt.Println(string(x))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	chain := shared.CreateChainConnection()
	err = chain.CreateAsset(&a)
	chain.CloseConnection()
	return "", err
}
