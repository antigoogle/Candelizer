package main

import (
	"fmt"
	"github.com/flower1024/Candelizer/bfx"
	"github.com/flower1024/Candelizer/candles"
	"github.com/flower1024/Candelizer/config"
	"github.com/flower1024/Candelizer/indicators"
)

func main() {
	fmt.Printf("START Candelizer\n")
	fin := make(chan bool)
	conf := config.ReadConfig()

	for symbol, market := range conf.BFX {
		go func(s string, m config.Market) {
			i := indicators.Indicate(candles.CandleStore(candles.Candelize(bfx.Trades(s), m.CandleWidthMins*60), m.KeepCandleMins*60), m.Indicators)

			for {
				candle := <-i.Update
				fmt.Printf("RECV HotCandle: BFX-%s %s\n", symbol, candle)
			}

		}(symbol, market)
	}

	<-fin
	fmt.Printf("QUIT Candelizer")
}
