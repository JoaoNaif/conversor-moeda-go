package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed data.json
var ratesJSON []byte

type RatesData struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func LoadRates() (RatesData, error) {
	var data RatesData
	err := json.Unmarshal(ratesJSON, &data)
	return data, err
}
