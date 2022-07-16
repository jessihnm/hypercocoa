package shared

import (
	"encoding/json"
)

// Farmer
type Farmer struct {
	ID        string
	Surname   string
	Firstname string
	Address   Address
	Phone     string
	Email     string
}

// ToAsset generalizes the Farmer to an Asset
func (f *Farmer) ToAsset() *Asset {
	return &Asset{
		ID:        f.ID,
		AssetType: AssetTypeFarmer,
		Refs:      map[string]string{},
		Data: map[string]string{
			"addressNumber":  f.Address.Number,
			"addressStreet":  f.Address.Street,
			"addressCity":    f.Address.City,
			"addressZipCode": f.Address.ZipCode,
			"addressCountry": f.Address.Country,
			"firstname":      f.Firstname,
			"surname":        f.Surname,
			"phone":          f.Phone,
			"email":          f.Email,
		},
	}
}

// ToJson converts the Farmer to json
func (f *Farmer) ToJson() ([]byte, error) {
	return json.Marshal(f)
}

// ToJson converts the Farmer to human readable json
func (f *Farmer) ToPrettyJson() ([]byte, error) {
	return json.MarshalIndent(f, "", "  ")
}

// toCocoaBag converts an Asset to Farmer struct
func (a *Asset) ToFarmer() *Farmer {
	return &Farmer{
		ID: a.ID,
		Address: Address{
			Number:  a.Data["addressNumber"],
			Street:  a.Data["addressStreet"],
			City:    a.Data["addressCity"],
			ZipCode: a.Data["addressZipCode"],
			Country: a.Data["addressCountry"],
		},
		Surname:   a.Data["surname"],
		Firstname: a.Data["firstname"],
		Phone:     a.Data["phone"],
		Email:     a.Data["email"],
	}
}
