package models

type SymbolLookup struct {
	Description   string `json:"description,omitempty"`
	DisplaySymbol string `json:"displaySymbol"`
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
}
