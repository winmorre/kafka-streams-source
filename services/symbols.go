package services

import (
	"context"
	"encoding/json"
	"fmt"
	"source/models"
)

func getSymbols() {
	finnhubClient := getApiService()
	res, _, err := finnhubClient.SymbolSearch(context.Background()).Q("AAPL").Execute()
	if err != nil {
		panic(err)
	}
	r, _ := json.Marshal(res.Result)
	var results []models.SymbolLookup
	err = json.Unmarshal(r, &results)
	if err != nil {
		panic(err)
	}
	for _, i := range results {
		fmt.Printf("%v\n", i.Description)
	}
}
