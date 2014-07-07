package candles

import (
	"fmt"
	"github.com/flower1024/Candelizer/types"
	"strconv"
)

func CandleStore(candles types.CandleChannel, keep int) types.CS {
	ccount := keep / candles.Width
	fmt.Printf("INIT CandleStore keep: " + strconv.Itoa(keep/(60*60)) + "h  (" + strconv.Itoa(keep/(ccount)) + "candles) " + candles.Exchange + " " + candles.Symbol + "\n")

	update := make(chan types.Candle)
	cs := types.CS{Update: update, Data: make([]types.Candle, ccount)}
	for i := 0; i < ccount; i++ {
		cs.Data[i].Indicators = make(map[string]float64)
	}

	go func() {
		for {
			select {
			case candle := <-candles.Channel:
				if candle.Indicators == nil {
					candle.Indicators = make(map[string]float64)
				}
				cs.Lock()
				for i := 0; i < ccount-2; i++ {
					cs.Data[i+1] = cs.Data[i]
				}
				cs.Data[0] = candle
				cs.Unlock()
			case hotCandle := <-candles.HotChannel:
				if hotCandle.Indicators == nil {
					hotCandle.Indicators = make(map[string]float64)
				}
				update <- hotCandle
			}

		}
	}()

	return cs
}
