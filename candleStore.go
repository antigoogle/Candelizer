package main

import (
	"fmt"
	"strconv"
)

type CS struct {
	data   []Candle
	Update <-chan Candle
}

func (cs CS) Request() []Candle {
	return cs.data
}

func CandleStore(candles CandleChannel, keep int) CS {
	ccount := keep / candles.Width
	fmt.Printf("INIT CandleStore keep: " + strconv.Itoa(keep/(60*60)) + "h  (" + strconv.Itoa(keep/(ccount)) + "candles) " + candles.Exchange + " " + candles.Symbol + "\n")

	update := make(chan Candle)
	cs := CS{Update: update, data: make([]Candle, ccount)}

	go func() {
		for {
			select {
			case candle := <-candles.Channel:
				for i := 0; i < ccount-2; i++ {
					cs.data[i+1] = cs.data[i]
				}
				cs.data[0] = candle
			case hotCandle := <-candles.HotChannel:
				update <- hotCandle
			}

		}
	}()

	return cs
}
