package shared

import (
	"bytes"
	"encoding/json"
)

type AssetType int64

const (
	AssetTypeCocoaBag AssetType = iota
	AssetTypeChocolateBar
	AssetTypeFarmer
	AssetTypeFarm
	AssetTypeFactory
	AssetTypeVendor
)

func (at AssetType) String() string {
	switch at {
	case AssetTypeCocoaBag:
		return "cocoabag"
	case AssetTypeChocolateBar:
		return "chocolatebar"
	case AssetTypeFarmer:
		return "farmer"
	case AssetTypeFarm:
		return "farm"
	case AssetTypeFactory:
		return "factory"
	case AssetTypeVendor:
		return "vendor"
	}
	return "unknown"
}

var assetTypeToID = map[string]AssetType{
	"cocoabag":     AssetTypeCocoaBag,
	"chocolatebar": AssetTypeChocolateBar,
	"farmer":       AssetTypeFarmer,
	"farm":         AssetTypeFarm,
	"factory":      AssetTypeFactory,
	"vendor":       AssetTypeVendor,
}

func (at AssetType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(at.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (at *AssetType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*at = assetTypeToID[j]
	return nil
}
