package services

import (
	"context"
	"fmt"
	"time"
)

func getIpoCalendar() {
	client := getApiService()
	res, _, err := client.IpoCalendar(context.Background()).From("2022-10-24").To(time.Now().String()).Execute()

	if err != nil {
		panic(err)
	}
	r, _ := res.MarshalJSON()
	fmt.Printf("%v \n", string(r))
}
