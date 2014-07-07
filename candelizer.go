package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var config struct {
	BFX map[string]Market
}

type Market struct {
	CandleWidthMins int
	KeepCandleMins  int
}

func main() {
	fmt.Printf("INIT Candelizer")

	fin := make(chan bool)
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("%s", err)
	}
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	file.Close()
	for symbol, market := range config.BFX {
		go func(s string, m Market) {
			i := Indicate(CandleStore(Candelize(BFXTrades(s), m.CandleWidthMins*60), m.KeepCandleMins*60))

			for {
				candle := <-i.Update
				fmt.Printf("%s %s; ", s, strconv.FormatFloat(candle.Close, 'f', 8, 64))
				//fmt.Printf("RECV HotCandle: BFX-%s %s\n", symbol, candle)
			}

		}(symbol, market)
	}

	<-fin
	fmt.Printf("QUIT Candelizer")
}
