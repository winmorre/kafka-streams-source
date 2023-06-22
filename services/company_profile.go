package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func GetCompanyProfile() {
	client := getApiService()
	emitter, err := getEmitter(profileTopic)
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}

	fmt.Println("Server about to start emitting company profile")
	for _, s := range symbols {
		res, _, err := client.CompanyProfile2(context.Background()).Symbol(s).Execute()
		if err != nil {
			log.Fatalf("Error getting company profile:  %v", err)
		}
		//var profile models.CompanyProfile
		marshalData, _ := json.Marshal(res)
		//_ = json.Unmarshal(marshalData, &profile)
		err = emitter.EmitSync(s, marshalData)

		if err != nil {
			log.Fatalf("Error emitting company profile: %v", err)
		}
	}
}
