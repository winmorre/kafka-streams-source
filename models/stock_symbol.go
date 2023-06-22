package models

type StockSymbol struct {
	Currency      string `json:"currency"`
	Description   string `json:"description"`
	DisplaySymbol string `json:"displaySymbol"`
	FIGI          string `json:"figi"`
	ISIN          string `json:"isin,omitempty"`
	MIC           string `json:"mic"`
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
}
