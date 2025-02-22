package data

import "encoding/json"

type Company struct {
	Name        string          `json:"name"`
	Logo        string          `json:"logo"`
	URL         string          `json:"url"`
	StaffCount  json.RawMessage `json:"staffCountRange"`
	Headquarter json.RawMessage `json:"headquarter"`
}
