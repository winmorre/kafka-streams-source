package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/lovoo/goka"
	"log"
	"strconv"
)

func GetTrades() {

	//defer (*wg).Done()
	w, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf(tradesUrl, apiKey), nil)
	if err != nil {
		println("the error occurred")
		panic(err)

	}
	defer w.Close()

	for _, s := range symbols {
		msg, _ := json.Marshal(map[string]interface{}{"type": "subscribe", "symbol": s})
		w.WriteMessage(websocket.TextMessage, msg)
	}

	var msg map[string]interface{}
	var emitter *goka.Emitter
	emitter, err = getEmitter(tradesTopic)

	if err != nil {
		log.Fatalf("Error creating emitter %v", err)
	}

	k := 1
	fmt.Println("Server starting to emit trades")
	for {
		err := w.ReadJSON(&msg)

		if err != nil {
			panic(err)
		}
		if err != nil {
			fmt.Printf("\n Error trade %v", err.Error())
		}
		b, _ := json.Marshal(msg["data"])

		err = emitter.EmitSync(strconv.Itoa(k), b)
		k += 1
		if err != nil {
			log.Fatalf("error emitting trade data %v", err)
		}

	}
}
