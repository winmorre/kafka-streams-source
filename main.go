package main

import (
	"source/services"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		services.GetTrades()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		services.GetCompanyBasicFinancials()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		services.GetMarketNews()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		services.GetCompanyProfile()
	}()
	wg.Wait()
}

// join company trade data with company profile, and also with company news
// get the company that's trading best
