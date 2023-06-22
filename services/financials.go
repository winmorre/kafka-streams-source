package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func GetCompanyBasicFinancials() {
	client := getApiService()

	emitter, err := getEmitter(financialsTopic)

	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}

	fmt.Println("Server about to emit company basic financials data")
	for _, s := range symbols {
		res, _, err := client.CompanyBasicFinancials(context.Background()).Symbol(s).Metric("all").Execute()
		if err != nil {
			log.Fatalf("")
		}
		marshalSeries, _ := json.Marshal(res.Series)
		err = emitter.EmitSync(s, marshalSeries)
		if err != nil {
			log.Fatalf("error emitting to %v", financialsTopic)
		}
	}

}
