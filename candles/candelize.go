package candles

import (
	"fmt"
	"github.com/flower1024/Candelizer/types"
	"strconv"
)

func Candelize(trades types.TradeChannel, width int) types.CandleChannel {
	fmt.Printf("INIT Candelizer width " + strconv.Itoa(width/60) + "min for " + trades.Exchange + " " + trades.Symbol + "\n")
	channel := make(chan types.Candle)
	hotchannel := make(chan types.Candle)
	cc := types.CandleChannel{Exchange: trades.Exchange, Symbol: trades.Symbol, Width: width, Channel: channel, HotChannel: hotchannel}

	go func() {
		current := types.Candle{}
		for {
			trade := <-trades.Channel
			candleDate := (trade.Date / width) * width

			if current.Date == candleDate { // update current candle
				if trade.Price < current.Low {
					current.Low = trade.Price
				}
				if trade.Price > current.High {
					current.High = trade.Price
				}
				current.Close = trade.Price
				current.Volume += trade.Amount

				hotchannel <- current
			} else if current.Date != 0 {
				// close current candle
				channel <- current

				// missing candles
				for d := current.Date + width; d < candleDate; d += width {
					channel <- types.Candle{Date: d, Open: current.Close, Low: current.Close, High: current.Close, Close: current.Close, Volume: 0}
				}
				// add new for current trade
				current = types.Candle{}
			}
			if current.Date == 0 { // old candle closed or first data
				current.Date = candleDate
				current.Open = trade.Price
				current.Low = trade.Price
				current.High = trade.Price
				current.Close = trade.Price
				current.Volume = trade.Amount
				hotchannel <- current
			}

		}
	}()
	return cc
}
