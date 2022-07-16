package shared

import (
	"encoding/json"
	"strconv"
)

type Address struct {
	Number  string
	Street  string
	City    string
	ZipCode string
	Country string
}

// Farm
type Farm struct {
	ID              string
	Name            string
	Address         Address
	Size            int
	FarmerReference string
}

// ToAsset generalizes the Farm to an Asset
func (f *Farm) ToAsset() *Asset {
	return &Asset{
		ID:        f.ID,
		AssetType: AssetTypeFarm,
		Refs: map[string]string{
			"farmer": f.FarmerReference,
		},
		Data: map[string]string{
			"addressNumber":  f.Address.Number,
			"addressStreet":  f.Address.Street,
			"addressCity":    f.Address.City,
			"addressZipCode": f.Address.ZipCode,
			"addressCountry": f.Address.Country,
			"size":           strconv.Itoa(f.Size),
			"name":           f.Name,
		},
	}
}

// ToJson converts the Farm to json
func (f *Farm) ToJson() ([]byte, error) {
	return json.Marshal(f)
}

// ToJson converts the Farm to human readable json
func (f *Farm) ToPrettyJson() ([]byte, error) {
	return json.MarshalIndent(f, "", "  ")
}

// toCocoaBag converts an Asset to Farm struct
func (a *Asset) ToFarm() *Farm {
	size, _ := strconv.Atoi(a.Data["size"])
	return &Farm{
		ID: a.ID,
		Address: Address{
			Number:  a.Data["addressNumber"],
			Street:  a.Data["addressStreet"],
			City:    a.Data["addressCity"],
			ZipCode: a.Data["addressZipCode"],
			Country: a.Data["addressCountry"],
		},
		FarmerReference: a.Refs["farmer"],
		Name:            a.Data["name"],
		Size:            size,
	}
}
