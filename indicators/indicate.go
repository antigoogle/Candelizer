package indicators

import (
	"fmt"
	"github.com/flower1024/Candelizer/types"
)

func Indicate(cs types.CS, indicators []map[string]interface{}) Indicators {
	fmt.Printf("INIT Indicators\n")

	update := make(chan types.Candle)
	idc := Indicators{Update: update, cs: cs}

	go func() {
		for {
			hotCandle := <-idc.cs.Update

			for i := 0; i < len(indicators); i++ {
				idc.Run(indicators[i], hotCandle)
			}

			update <- hotCandle
		}
	}()

	return idc
}
