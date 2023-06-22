package services

import (
	"context"
	"encoding/json"
	"log"
)

func GetMarketNews() {
	client := getApiService()
	emitter, err := getEmitter(marketNewsTopic)
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	newsCategory := []string{"general", "forex", "crypto", "merger"}

	for _, nc := range newsCategory {
		res, _, err := client.MarketNews(context.Background()).Category(nc).Execute()
		if err != nil {
			log.Fatalf("Error getting market news: %v", err)
		}
		marshalNews, _ := json.Marshal(res)
		err = emitter.EmitSync(nc+"-news", marshalNews)
		if err != nil {
			log.Fatalf("error emitting market news: %v", err)
		}
	}
}
