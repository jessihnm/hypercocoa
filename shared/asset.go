package shared

import (
	"encoding/json"
	"fmt"
)

type Asset struct {
	ID        string
	AssetType AssetType
	Refs      map[string]string
	Data      map[string]string
}

func (a *Asset) ToJson() ([]byte, error) {
	return json.Marshal(a)
}

func (a *Asset) ToPrettyJson() ([]byte, error) {
	return json.MarshalIndent(a, "", "  ")
}

type assetUnmarshalHelper struct {
	ID        string
	AssetType AssetType
	Refs      string
	Data      string
}

func (a *assetUnmarshalHelper) ToPrettyJson() ([]byte, error) {
	return json.MarshalIndent(a, "", "  ")
}

// UnmarshalJSON customizes unmarshalling a helper class to handle escaped responsed from the ledger
// Refs and Data are escaped by the ledger, so we need to unescape it
// For cases where it is not escaped, e.g. Rest API, we read as is
func (a *assetUnmarshalHelper) UnmarshalJSON(b []byte) error {
	var i map[string]interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		return err
	}
	a.ID = i["ID"].(string)
	if _, ok := i["Refs"].(string); ok {
		a.Refs = i["Refs"].(string)
	} else {
		tmp, err := json.Marshal(i["Refs"])
		if err != nil {
			fmt.Println(err)
			return err
		}
		a.Refs = string(tmp)
	}

	if o, ok := i["Data"].(string); ok {
		a.Data = o
	} else {
		tmp, err := json.Marshal(i["Data"])
		if err != nil {
			fmt.Println(err)
			return err
		}
		a.Data = string(tmp)
	}
	a.AssetType = assetTypeToID[i["AssetType"].(string)]

	return nil
}

// UnmarshalJSON unmarshals the Asset using the helper class
func (a *Asset) UnmarshalJSON(b []byte) error {
	ah := new(assetUnmarshalHelper)
	err := json.Unmarshal(b, &ah)
	if err != nil {
		return err
	}
	a.ID = ah.ID
	a.AssetType = ah.AssetType
	err = json.Unmarshal([]byte(ah.Data), &a.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(ah.Refs), &a.Refs)
}
