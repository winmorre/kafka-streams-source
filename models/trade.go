package models

type Trade struct {
	Symbol    string   `json:"s"`
	Price     float64  `json:"p"`
	Timestamp float64  `json:"t"`
	Volume    float64  `json:"v"`
	C         []string `json:"c,omitempty"`
}
