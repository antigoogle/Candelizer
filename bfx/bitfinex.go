package bfx

import (
	"encoding/json"
	"fmt"
	"github.com/flower1024/Candelizer/types"
	"net/http"
	"strconv"
	"time"
)

func Trades(symbol string) types.TradeChannel {
	url := "https://api.bitfinex.com/v1/trades/" + symbol
	channel := make(chan types.Trade)
	tc := types.TradeChannel{Exchange: "BFX", Symbol: symbol, Channel: channel}

	lastTS := 0
	lastTID := 0
	fmt.Printf("INIT BFX.Trades %s\n", url)

	go func() {
		for {
			timer := time.NewTimer(5 * time.Second)
			reqUrl := url
			if lastTS != 0 {
				reqUrl += "?timestamp=" + strconv.Itoa(lastTS)
			}
			fmt.Printf("HTTP %s\n", reqUrl)
			response, err := http.Get(reqUrl)
			if err != nil {
				fmt.Printf("%s", err)
				continue
			} else {
				func() {
					defer response.Body.Close()

					var value []map[string]interface{}
					err = json.NewDecoder(response.Body).Decode(&value)
					if err != nil {
						fmt.Printf("%s", err)
						return
					}

					lTS := lastTS
					lTID := lastTID
					foundTID := false

					for i := len(value) - 1; i >= 0; i-- {
						timestamp := int(value[i]["timestamp"].(float64))
						tid := int(value[i]["tid"].(float64))
						if lTS != 0 && foundTID == false && lTS == timestamp {
							foundTID = tid == lTID
							continue
						}

						price, err := strconv.ParseFloat(value[i]["price"].(string), 64)
						if err != nil {
							continue
						}
						amount, err := strconv.ParseFloat(value[i]["amount"].(string), 64)
						if err != nil {
							continue
						}

						trade := types.Trade{Date: timestamp, Price: price, Amount: amount}

						channel <- trade
					}

					if len(value) > 0 {
						lastTS = int(value[0]["timestamp"].(float64))
						lastTID = int(value[0]["tid"].(float64))
					}

				}()
			}

			<-timer.C
		}
	}()

	return tc
}
