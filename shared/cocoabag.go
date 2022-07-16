package shared

import (
	"encoding/json"
	"strconv"
)

// CocoaBag
type CocoaBag struct {
	ID               string
	Type             string
	Weight           int
	FarmReference    string
	FactoryReference string `json:",omitempty"`
	Price            int
	Currency         string
}

// ToAsset generalizes the Cocoabag to an Asset
func (c *CocoaBag) ToAsset() *Asset {
	a := &Asset{
		ID:        c.ID,
		AssetType: AssetTypeCocoaBag,
		Refs: map[string]string{
			"farm": c.FarmReference,
		},
		Data: map[string]string{
			"price":    strconv.Itoa(c.Price),
			"currency": c.Currency,
			"weight":   strconv.Itoa(c.Weight),
			"type":     c.Type,
		},
	}
	if c.FactoryReference != "" {
		a.Refs["factory"] = c.FactoryReference
	}
	return a
}

// ToJson converts the CocoaBag to json
func (c *CocoaBag) ToJson() ([]byte, error) {
	return json.Marshal(c)
}

// ToJson converts the CocoaBag to human readable json
func (c *CocoaBag) ToPrettyJson() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

// toCocoaBag converts an Asset to CocoaBag struct
func (a *Asset) ToCocoaBag() *CocoaBag {
	weight, _ := strconv.Atoi(a.Data["weight"])
	price, _ := strconv.Atoi(a.Data["price"])
	return &CocoaBag{
		ID:            a.ID,
		Weight:        weight,
		FarmReference: a.Refs["farm"],
		Price:         price,
		Currency:      a.Data["price"],
		Type:          a.Data["type"],
	}
}
