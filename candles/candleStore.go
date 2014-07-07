package candles

import (
	"fmt"
	"github.com/flower1024/Candelizer/types"
	"strconv"
)

type CS struct {
	data   []types.Candle
	Update <-chan types.Candle
}

func (cs CS) Request() []types.Candle {
	return cs.data
}

func CandleStore(candles types.CandleChannel, keep int) CS {
	ccount := keep / candles.Width
	fmt.Printf("INIT CandleStore keep: " + strconv.Itoa(keep/(60*60)) + "h  (" + strconv.Itoa(keep/(ccount)) + "candles) " + candles.Exchange + " " + candles.Symbol + "\n")

	update := make(chan types.Candle)
	cs := CS{Update: update, data: make([]types.Candle, ccount)}

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
