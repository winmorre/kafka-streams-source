package models

type CompanyProfile struct {
	Country              string  `json:"country"`
	Currency             string  `json:"currency"`
	Exchange             string  `json:"exchange"`
	IPO                  string  `json:"ipo"`
	MarketCapitalization int64   `json:"marketCapitalization"`
	Name                 string  `json:"name"`
	Phone                string  `json:"phone"`
	ShareOutstanding     float64 `json:"shareOutstanding"`
	Ticker               string  `json:"ticker"`
	WebUrl               string  `json:"weburl"`
	Logo                 string  `json:"logo"`
	FinnhubIndustry      string  `json:"finnhubIndustry"`
}
