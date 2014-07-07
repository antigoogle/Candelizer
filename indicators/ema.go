package indicators

import (
	"github.com/flower1024/Candelizer/types"
)

func EMA(ema float64, field string, candles []types.Candle, hotCandle types.Candle) {

	if len(candles) > 0 {
		var k float64
		var last_ema float64
		for i := len(candles) - 1; i >= 0; i-- {
			k = float64(2) / (ema + float64(1))
			last_ema = candles[i].Close

			if i < len(candles)-1 {
				last_ema = candles[i+1].Indicators[field]
			}

			candles[i].Indicators[field] = candles[i].Close*k + last_ema*(1-k)
		}

		hotCandle.Indicators[field] = candles[0].Close*k + last_ema*(1-k)
	}
}
