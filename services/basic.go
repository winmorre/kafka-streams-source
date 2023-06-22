package services

import (
	"github.com/Finnhub-Stock-API/finnhub-go/v2"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

const (
	apiKey    string = "ci6l4r9r01qhmud1lfj0ci6l4r9r01qhmud1lfjg"
	tradesUrl string = "wss://ws.finnhub.io?token=%v"
)

var (
	brokers                     = []string{"localhost:29092"}
	profileTopic    goka.Stream = "company-profile"
	financialsTopic goka.Stream = "basic-financials"
	tradesTopic     goka.Stream = "trades"
	marketNewsTopic goka.Stream = "market-news"
	symbols                     = []string{"AAPL", "AMZN", "BINANCE:BTCUSDT", "TSLA", "GNRC", "SEDG", "ENPH", "WBD", "PYPL", "WBD", "GOOGL", "MSFT", "NVDA", "META", "AMD", "NFLX", "ORCL", "BA", "DICE"}
)

func getApiService() *finnhub.DefaultApiService {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", apiKey)
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi
	return finnhubClient
}

func getEmitter(topic goka.Stream) (*goka.Emitter, error) {
	emitter, err := goka.NewEmitter(brokers, topic, new(codec.Bytes))
	return emitter, err
}
